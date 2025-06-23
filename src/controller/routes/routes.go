package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/controller"
	"github.com/kayquesza/gocrud-auth-api/src/model"
)

func InitRoutes( // Função que recebe um grupo de rotas do Gin e um controlador de usuário
	r *gin.RouterGroup, // Grupo de rotas do Gin
	userController controller.UserControllerInterface, // Controlador de usuário
) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)          // Rota para buscar um usuário por ID
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail) // Rota para buscar um usuário por email
	r.POST("/createUser", userController.CreateUser)                                                 // Rota para criar um usuário
	r.PUT("/updateUser/:userId", userController.UpdateUser)                                          // Rota para atualizar um usuário
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)                                       // Rota para deletar um usuário

	r.POST("/login", userController.LoginUser) // Rota para fazer login
}
