package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserById/:UserId", userController.FindUserByID)
	r.GET("/getUserByEmail/:UserEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:UserId", userController.UpdateUser)
	r.DELETE("/deleteUser/:UserId", userController.DeleteUser)
}
