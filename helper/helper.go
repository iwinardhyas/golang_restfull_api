package helper

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Db() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().
		ApplyURI(""))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// datapeople := client.Database("stream").Collection("user")
	// cursor, err := datapeople.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var email []bson.M
	// if err = cursor.All(ctx, &email); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(email)
	return client, ctx
}
