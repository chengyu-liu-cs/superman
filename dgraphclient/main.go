

func main() {
    //logger.Init()
    //Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
    // connect
    dg, conn := dgraphclient.NewClient("localhost", "9080")
    defer dgraphclient.Close(conn)

    schema := `
        name: string @index(exact) .
        age: int .
    `
    //alter database schema
    dgraphclient.AlterSchema(dg, schema)

    var event []entity.Event
    event = read_csv("../../data/shogun_content.csv", event)
    for i :=0; i < len(event); i++ {
        logger.Info().Println(i)
        dgraphclient.InsertAnEdge(dg, event[i].User, event[i].Article)
    }
    logger.Info().Println("Hello " + "world")
}
package main

import (
    "context"
    "log"
    "os"
    "io"
    "bufio"
    "encoding/csv"
    "dgraphclient"
    "logger"
    "entity"

    "github.com/dgraph-io/dgraph/client"
)
