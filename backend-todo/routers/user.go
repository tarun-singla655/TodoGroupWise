package routers

import (
	"React-golang/controllers"

	"github.com/julienschmidt/httprouter"
)

func InitUserRouter(r *httprouter.Router, uc *controllers.UserController) {
	r.POST("/user/signup", uc.SignUp)
	r.POST("/user/login", uc.Login)
	r.GET("/users", uc.GetAllUsers)
}
