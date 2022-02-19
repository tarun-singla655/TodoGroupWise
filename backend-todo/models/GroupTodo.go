package models

import (
	"time"
	//  "React-golang/models"
	"gopkg.in/mgo.v2/bson"
)

type TodoAssigned struct {
	UserFor  bson.ObjectId `json:"userfor" bson:"userfor" `
	TodoUser bson.ObjectId `json:"todo_user" bson:"todo_user" `
}
type GroupTodo struct {
	Id        bson.ObjectId  `json:"id" bson:"_id"  `
	Completed bool           `json:"completed" bson:"completed"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at"  `
	UpdateAt  time.Time      `json:"updated_at" bson:"updated_at"  `
	Createdby bson.ObjectId  `json:"created_by" bson:"created_by"  `
	Todos     []TodoAssigned `json:"todos" bson:"todos" `
}
