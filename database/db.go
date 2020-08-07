package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

//database name
const DBNAME = "frsc"

// collection name
const COLLECTION = "bus"

// collection object/instance
var Collection *mongo.Collection

func init() {
	MONGODB := os.Getenv("MONGODB")

	// Set the client options
	clientOptions := options.Client().ApplyURI(MONGODB)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongo successfully")
	Collection = client.Database(DBNAME).Collection(COLLECTION)

	fmt.Println("connection instance created")
}
