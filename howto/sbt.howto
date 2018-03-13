#############  SBT  ############
################################

# SBT folder and file structure:
$ find .
.
./build.sbt
./src
./src/main
./src/main/scala
./src/main/scala/SimpleApp.scala

# build.sbt


# Package a jar containing your application
$ sbt package

# Use spark-submit to run your application
$ YOUR_SPARK_HOME/bin/spark-submit \
                    #--class "SimpleApp" \
                    --master local[4] \
                    target/scala-2.11/simple-project_2.11-1.0.jar
