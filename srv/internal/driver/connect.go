package driver

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB name
const dbName = "weigh-me-cluster"

// Collection name
const collectionName = "weight"

// Collection object
var collection *mongo.Collection

// Connect will connect to the mongo database
func Connect() (*mongo.Collection, error) {
	godotenv.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectString := fmt.Sprintf("mongodb+srv://%s:%s@%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_CLUSTER"))

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectString))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Print("Connected to MongoDB")
	collection = client.Database(dbName).Collection(collectionName)

	return collection, nil
}
