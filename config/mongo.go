package config

import (
	"context"
	"fmt"
	"time"

	"github.com/nandonyata/Stray-Fedding/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase() *mongo.Database {
	uri := "mongodb://localhost:27017"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	exception.PanicIfNeeded(err)

	defer func() {
		err = client.Disconnect(context.TODO())
		exception.PanicIfNeeded(err)
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	err = client.Database("golang").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	exception.PanicIfNeeded(err)

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	database := client.Database("golang")
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
