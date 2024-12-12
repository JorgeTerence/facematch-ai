package main

import (
	"context"
	"env"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_PATTERN = "mongodb+srv://%s:%s@cluster0.3t4vm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

type DB struct {
	username string
	password string
	client   *mongo.Client
	env      env.Environment
}

func (db DB) Connect() error {
	connectionString := fmt.Sprintf(MONGODB_PATTERN, db.username, db.password)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err = client.Database(string(db.env)).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return err
	}

	return nil
}
