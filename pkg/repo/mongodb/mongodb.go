package mongodb

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var db *mongo.Database

func InitDatabase() error {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
		return err
	}

	mongoHost := os.Getenv("MONGODB_HOST")
	mongoPort := os.Getenv("MONGODB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct the connection string
	connectionString := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	clientOptions := options.Client().ApplyURI(connectionString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	db = client.Database(dbName)

	return nil
}

// GetDB returns the MongoDB database instance.
func GetDB() *mongo.Database {
	return db
}
