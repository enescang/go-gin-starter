package db

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var err error
var ctx context.Context
var mongoSingleton sync.Once

func Init() (*mongo.Database, error) {
	mongoSingleton.Do(func() {
		client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		err = client.Ping(ctx, readpref.Primary())
	})
	return client.Database(os.Getenv("DB_NAME")), nil
}

func Disconnect() {
	client.Disconnect(ctx)
}
