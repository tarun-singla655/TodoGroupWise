package main

import (
	"React-golang/controllers"
	"React-golang/routers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	routers.InitTodoRouter(r, uc)
	routers.InitUserRouter(r, uc)
	routers.InitGroupTodoRouter(r, uc)
	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	fmt.Println("hi i am connected")
	return s
}
