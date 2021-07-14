package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/iwinardhyas/restapi/helper"
	model "github.com/iwinardhyas/restapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Db() {
	client, err := mongo.NewClient(options.Client().
		ApplyURI("mongodb+srv://stream:abcd12345@cluster0.3515n.mongodb.net/Cluster0?retryWrites=true&w=majority"))
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
}

// var client *mongo.Client

func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	client, ctx := helper.Db()
	datapeople := client.Database("stream").Collection("user")
	cursor, err := datapeople.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

func Parsing_json(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	b, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(b, &person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	email := person.Email
	w.Write([]byte(email))
	fmt.Println(email)
}

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var person model.Person
	json.NewDecoder(r.Body).Decode(&person)
	// collection := client.Database("stream").Collection("user")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// result, _ := collection.InsertOne(ctx, person)
	// json.NewEncoder(w).Encode(result)
}

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	// var pople []model.Person
	collection := client.Database("stream").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	curser, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer curser.Close(ctx)

}
