package controllers

import (
	"React-golang/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

func (uc UserController) CheckDuplicate(email string) bool {
	count, _ := uc.session.DB("TodoApp").C("users").Find(bson.M{"email": email}).Count()
	if count > 0 {
		return true
	}
	return false
}

func (uc UserController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// var user models.User
	var login models.Userlogin
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("Network Error %s", w)
		panic(err)
	}

	if !uc.CheckDuplicate(login.Email) {
		fmt.Println("email does not  exist")
		w.WriteHeader(409)
		// fmt.Errorf("email does not exist %v", w)
	}
	var user models.User
	err = uc.session.DB("TodoApp").C("users").Find(bson.M{"email": login.Email}).One(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Password != login.Password {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userj, err := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userj)

}

func (uc UserController) SignUp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if uc.CheckDuplicate(user.Email) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Id = bson.NewObjectId()
	user.Todos = []models.Todo{}
	uc.session.DB("TodoApp").C("users").Insert(user)
	userj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userj)
}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	users := []models.User{}
	if err := uc.session.DB("TodoApp").C("users").Find(nil).All(&users); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usersj, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", usersj)

}
