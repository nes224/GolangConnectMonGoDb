package main


import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
// Set up the MongoDB connection options
func main(){
	clientOptions := options.Client().ApplyURI("mongodb+srv://<user>:<password>@cluster0.sclsoat.mongodb.net/?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return
	}

	fmt.Println("Connected to MongoDB!")
}