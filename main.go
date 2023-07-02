package main


import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
// Set up the MongoDB connection options
func main(){
  // Use the SetServerAPIOptions() method to set the Stable API version to 1
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI("mongodb+srv://<user>:<password>@cluster0.sclsoat.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
  // Create a new client and connect to the server
  client, err := mongo.Connect(context.TODO(), opts)
  if err != nil {
    panic(err)
  }
  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()
  // Check the connection
  err = client.Ping(context.TODO(),nil)
  if err != nil{
	fmt.Println("Failed to ping MongoDB:",err)
	return
  }
  // Send a ping to confirm a successful connection

  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}