package configs

import (
	"context"
	"os"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)
const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb+srv://%s:%s@%s"
)

func ConnectDB() *mongo.Client {
	// Safety to check if we can load the env credentials
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Extract env variables from .env file
	username        := os.Getenv("MONGODB_USERNAME")
	password        := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	// build the MonogoDB URI with the .env data
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	// Instanciate the client with the mongoDB URI
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	// Safety check
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collection string) *mongo.Collection {
	return client.Database("usersCollection").Collection(collection)
}