package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Userlogin struct {
	Email    string `json:"email" bson:"email" `
	Password string `json:"password" bson:"password"`
}

type UserName struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
}

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"  `
	FirstName string        `json:"firstname" bson:"firstname"  `
	LastName  string        `json:"lastname" bson:"lastname"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password" `
	Todos     []Todo        `json:"todos" bson:"todos"`
}
