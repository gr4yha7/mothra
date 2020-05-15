package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gr4yha7/mothra/models"
	"github.com/gr4yha7/mothra/utils"

	// "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	var host, port, dbName string
	host = utils.GetEnvVar("MONGODB_HOST")
	port = utils.GetEnvVar("MONGODB_PORT")
	dbName = utils.GetEnvVar("MONGODB_DATABASE")
	// dbName = "spaceDB"

	mongodbURI := fmt.Sprintf("mongodb://%v:%v", host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Failed to ping Mongo DB Server", err)
	}
	db := client.Database(dbName)
	// log.Println("database: ", db.Name())

	// filter := bson.D{}
	// databases, err := client.ListDatabases(ctx, filter, &options.ListDatabasesOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(databases)

	userCollection := db.Collection("users")
	user := &models.User{
		Email:    "jamal@gmail.com",
		Password: "password123",
		Name:     "Jamal Ekpenyong",
	}
	hashPwd, hashError := utils.HashPassword(user.Password)
	if hashError != nil {
		log.Fatal(hashError)
	}
	user.Password = hashPwd

	insertResult, err := createUser(ctx, userCollection, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(insertResult.InsertedID)
}

func createUser(ctx context.Context, col *mongo.Collection, user *models.User) (*mongo.InsertOneResult, error) {
	result, err := col.InsertOne(ctx, &user)
	if err != nil {
		return result, err
	}
	return result, nil
}
