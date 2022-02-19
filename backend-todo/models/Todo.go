package models

import (
	"time"
	//  "React-golang/models"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id        bson.ObjectId `json:"id" bson:"_id"  `
	Text      string        `json:"text" bson:"text"  `
	Completed bool          `json:"completed" bson:"completed"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"  `
	UpdateAt  time.Time     `json:"updated_at" bson:"updated_at"  `
}
