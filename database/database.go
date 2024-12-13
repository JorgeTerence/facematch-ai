package database

import (
	"context"
	"env"
	"fmt"
	"os"
	"platform"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if err = client.Database(string(db.env)).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return err
	}

	return nil
}

func FromEnv() DB {
	return DB{os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"), nil, env.FromString(os.Getenv("ENV"))}
}

func (db DB) InsertAccount(account *platform.Account) error {
	database := db.client.Database(string(db.env))
	collection := database.Collection("accounts")
	res, err := collection.InsertOne(context.Background(), bson.D{{"username", account.Username}, {"platform_id", account.PlatformId}})
	if err != nil {
		return err
	}

	account.InternalId = res.InsertedID
	return nil
}

func (db DB) addAccountEmbedding(id primitive.ObjectID, embedding []rune) error
