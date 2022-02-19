package controllers

// package models
import (
	"React-golang/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// setupCorsResponse(&w, r)
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	var todo models.Todo
	uc.session.DB("TodoApp").C("todos").FindId(oid).One(&todo)
	todoj, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", todoj)
}

func (uc UserController) GetAllTodos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	todos := []models.Todo{}
	if err := uc.session.DB("TodoApp").C("todos").Find(nil).All(&todos); err != nil {
		fmt.Println("Sorry there is err", err)
	}
	todosj, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("Sorry there is err", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", todosj)
}

func (uc UserController) AddTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// setupCorsResponse(&w, r)

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	todo.UpdateAt = time.Now()
	todo.CreatedAt = time.Now()
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.Id = bson.NewObjectId()
	fmt.Println(todo)
	uc.session.DB("TodoApp").C("todos").Insert(todo)
	todoj, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", todoj)
}

func (uc UserController) DeleteUserTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uc.session.DB("TodoApp").C("todos").RemoveId(todo.Id)

	if err != nil {
		fmt.Println(err)
		return
	}

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("TodoApp").C("users").UpdateId(oid, bson.M{"$pull": bson.M{"todos": bson.M{"_id": todo.Id}}}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Deleted")

}
func (uc UserController) UpdateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.UpdateAt = time.Now()
	if err := uc.session.DB("TodoApp").C("todos").UpdateId(todo.Id, &todo); err != nil {
		fmt.Println("Sorry error is ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		fmt.Print("sorry sir error is its not bson")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("TodoApp").C("users").Update(bson.M{"_id": oid, "todos._id": todo.Id}, bson.M{"$set": bson.M{"todos.$": todo}}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", "Updated")

}

func (uc UserController) AddUserTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	todo.UpdateAt = time.Now()
	todo.CreatedAt = time.Now()
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.Id = bson.NewObjectId()
	uc.session.DB("TodoApp").C("todos").Insert(todo)
	todoj, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("TodoApp").C("users").UpdateId(oid, bson.M{"$push": bson.M{"todos": todo}}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", todoj)

}
