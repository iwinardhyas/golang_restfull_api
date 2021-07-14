package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/iwinardhyas/restapi/controllers"
	"github.com/iwinardhyas/restapi/helper"
	model "github.com/iwinardhyas/restapi/models"
)

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person model.Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	// collection := client.Database("stream").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// result, _ := collection.InsertOne(ctx, person)
	// json.NewEncoder(response).Encode(result)
	fmt.Print(person)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func handlerequest() {
	r := mux.NewRouter()
	helper.Db()
	r.HandleFunc("/", homepage)
	r.HandleFunc("/article/", controller.Parsing_json).Methods("POST")
	r.HandleFunc("/parsing_json", controller.Parsing_json).Methods("POST")
	log.Fatal(http.ListenAndServe(":1000", r))
}

func main() {
	handlerequest()
	// client, err := mongo.NewClient(options.Client().
	// 	ApplyURI("mongodb+srv://stream:abcd12345@cluster0.3515n.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
