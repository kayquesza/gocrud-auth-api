package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:UserId", controller.FindUserById)
	r.GET("/getUserByEmail/:UserEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:UserId", controller.UpdateUser)
	r.DELETE("/deleteUser/:UserId", controller.DeleteUser)
}
