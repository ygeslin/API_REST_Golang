package routes

import (
	controller "dep/API_REST_Golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.POST("add/users", controller.CreateUser)
	router.POST("login", controller.Login)
	router.DELETE("/delete/user/:id", controller.DeleteUser)
	router.GET("/user/:id", controller.GetAUser)
	router.GET("/users/list", controller.GetUserList)
	router.PATCH("/user/:id", controller.UpdateAUser)
}