package routers

import (
	"React-golang/controllers"

	"github.com/julienschmidt/httprouter"
)

func InitGroupTodoRouter(r *httprouter.Router, uc *controllers.UserController) {
	r.POST("/grouptodo/addtodo", uc.AddGroupTodo)
}
