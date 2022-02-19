package controllers

// package models
import (
	"React-golang/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (uc UserController) AddGroupTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// setupCorsResponse(&w, r)

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var todo models.GroupTodo

	err := json.NewDecoder(r.Body).Decode(&todo)
	todo.UpdateAt = time.Now()
	todo.CreatedAt = time.Now()
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bsonid := bson.ObjectIdHex(id)
	todo.Createdby = bsonid
	todo.Id = bson.NewObjectId()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.Id = bson.NewObjectId()
	fmt.Println(todo)
	uc.session.DB("TodoApp").C("usertodo").Insert(todo)
	todoj, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", todoj)
}
