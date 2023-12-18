package config

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(configuration IConfig) (context.Context, *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongoPoolMin, mpmErr := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	if mpmErr != nil {
		panic("mongoPoolMin unknown")
	}

	mongoPoolMax, poolMaxErr := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	if poolMaxErr != nil {
		panic("poolMaxErr unknown")
	}

	mongoMaxIdleTime, maxIdleTimeErr := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	if maxIdleTimeErr != nil {
		panic("maxIdleTimeErr unknown")
	}

	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, clientErr := mongo.NewClient(option)
	if clientErr != nil {
		panic("Failed to connect to database!")
	}

	if err := client.Connect(ctx); err != nil {
		panic("Failed to connect to database!")
	}

	log.Println("Connected to MongoDB success")

	return ctx, client.Database(configuration.Get("MONGO_DATABASE"))
}
