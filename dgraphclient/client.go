package dgraphclient

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "logger"
    "entity"

    "github.com/dgraph-io/dgraph/client"
    "github.com/dgraph-io/dgraph/protos/api"
    "google.golang.org/grpc"
)

//Alter database schema
func AlterSchema(dg *client.Dgraph, schema string){
    op := &api.Operation{}
    op.Schema = schema
    err := dg.Alter(context.Background(), op)
    if err != nil {
        fmt.Println("Altering schema failed")
        log.Fatal(err)
    }
}

func CheckNodeExistence(dg *client.Dgraph, node_id string) string{
    logger.Info().Printf("Checking whether the node (%s) exists", node_id)
    q := `query q($node_id: string){
                        node_uid(func: eq(name, $node_id)){
                            uid
                        }
                    }`
    parameters := make(map[string]string)
    parameters["$node_id"] = node_id
    var parser map[string][]map[string]string
    res := Query_with_parameters(dg, q, parameters)
    logger.Info().Println(string(res))
    err := json.Unmarshal(res, &parser)
    if err != nil {
        log.Fatal(err)
    }
    count := len(parser["node_uid"])
    if(count > 1){
        log.Fatal("Fatal: Mapped to multiple users")
        return ""
    }else if (count == 1){
        return parser["node_uid"][0]["uid"]
    }else{
        log.Printf("The node (%s) does not exist", node_id)
        return ""
    }
}

// Close the connection in the end.
func Close(conn *grpc.ClientConn){
    conn.Close()
}

func InsertAnEdge(dg *client.Dgraph, source string, target string) {
    /* When insert an edge, should check
        1. whether two nodes exist in the graph
        2. whether an edge exists between two nodes
            a. if not exist, then insert
            b. if exist, then update information about the edge (TODO)
    */
    logger.Info().Printf("Inserting an edge between %s and %s", source, target)
    source_uid := CheckNodeExistence(dg, source)
    target_uid := CheckNodeExistence(dg, target)

    user :=entity.User{}
    if( source_uid == ""){
        logger.Info().Println("Source does not exist")
        user.Xid = source
    }else{
        logger.Info().Printf("Source exists (%s)", source_uid)
        user.Uid = source_uid
    }
    if( target_uid == ""){
        logger.Info().Println("Target does not exist")
        user.Visited = entity.Article{Xid: target}
    }else{
        logger.Info().Printf("Target exists (%s)", target_uid)
        user.Visited = entity.Article{Uid: target_uid}
    }
    if(source_uid != "" && target_uid != "" ){
        if(!IsEdgeExistence(dg, source_uid, target_uid)){
            logger.Info().Printf("Edge does not exist")
            mutate(dg, user)
        }
    }else{
        mutate(dg, user)
    }
}

func IsEdgeExistence(dg *client.Dgraph, source string, target string) bool{
    logger.Info().Printf("Checking whether edge exists between %s and %s", source, target)
    q := `{ 
            edge_exist(func: uid(%s)){  
                visited @filter(uid(%s)){
                    uid
                }
            }
        }`
    q = fmt.Sprintf(q, source, target)
    parameters := make(map[string]string)
    parameters["$source"] = source
    parameters["$target"] = target
    logger.Info().Println(parameters)
    var parser map[string][]map[string][]map[string]string
    res := Query(dg, q)
    logger.Info().Println(string(res))
    err := json.Unmarshal(res, &parser)
    if err != nil {
        logger.Error().Println(err)
        log.Fatal(err)
    }
    count := len(parser["edge_exist"][0]["visited"])
    if(count > 1){
        logger.Error().Println(parser["edge_exist"][0]["visited"])
        log.Fatal("Fatal: There are multiple edges")
        return false
    }else if (count == 1){
        return true
    }else{
        return false
    }
    return false
}

func mutate(dg *client.Dgraph, user entity.User) *api.Assigned{
    logger.Info().Println("Mutating")
    pb, err := json.Marshal(user)
    if err != nil {
        logger.Error().Println("Marshalling failed")
        log.Fatal(err)
    }
    logger.Info().Printf("Updating user %s", pb)
    mu := &api.Mutation{
        CommitNow: true,
    }

    mu.SetJson = pb
    assigned, err := dg.NewTxn().Mutate(context.Background(), mu)
    if err != nil {
        log.Fatal(err)
    }
    return assigned
}

// Create a new dgraph client 
func NewClient(host string, port string) (*client.Dgraph, *grpc.ClientConn) {
    // Dial a gRPC connection. The address to dial to can be configured when
    // setting up the dgraph cluster.
    conn, err := grpc.Dial(host + ":" + port, grpc.WithInsecure())
    if err != nil {
        fmt.Println("Connection failed")
        log.Fatal(err)
    }

    // A single client is thread safe for sharing with multiple go routines.
    dc := api.NewDgraphClient(conn)
    dg := client.NewDgraphClient(dc)
    return dg, conn
}

// execute a query
func Query(dg *client.Dgraph, query string)[]byte{
    txn := dg.NewTxn()
    defer txn.Discard(context.Background())
    resp, err := txn.Query(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }
    // resp is *api.Response object
    return resp.Json
}

// execute a query with parameters
func Query_with_parameters(dg *client.Dgraph, query string, parameters map[string]string)[]byte{
    txn := dg.NewTxn()
    defer txn.Discard(context.Background())
    resp, err := txn.QueryWithVars(context.Background(), query, parameters)
    if err != nil {
        log.Fatal(err)
    }
    // resp is *api.Response object
    return resp.Json
}

