package config

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(configuration IConfig) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	exception.PanicIfNeeded(err)

	mongoPoolMax, err := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	exception.PanicIfNeeded(err)

	mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	exception.PanicIfNeeded(err)

	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = client.Connect(ctx)
	if err != nil {
		panic("Failed to connect to database!")
	}
	exception.PanicIfNeeded(err)

	log.Println("Connected to MongoDB success")

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}
