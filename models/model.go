package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	Id       primitive.ObjectID `json: "ID"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
