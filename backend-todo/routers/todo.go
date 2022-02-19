package routers

import (
	"React-golang/controllers"

	"github.com/julienschmidt/httprouter"
)

func InitTodoRouter(r *httprouter.Router, uc *controllers.UserController) {
	r.GET("/todos", uc.GetAllTodos)
	r.POST("/todo", uc.AddTodo)
	r.GET("/todo/:id", uc.GetTodo)
	r.DELETE("/todo/:id", uc.DeleteUserTodo)
	r.PUT("/todo/:id", uc.UpdateTodo)
	r.POST("/user/todos/:id", uc.AddUserTodo)
}
