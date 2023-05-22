package Mongo

import (
	config "backend/config"
	Log "backend/pkg/helpers/log"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	Client *mongo.Client
}

// ConnectMongoDB opens a connection to MongoDB.
func ConnectMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo_Host))
	if Log.Err(err) {
		panic(err)
	}
	if err := client.Ping(context.TODO(), nil); Log.Err(err) {
		panic(err)
	}
	log.Println("connected mongoclient:", config.Mongo_Host)
	return client
}
