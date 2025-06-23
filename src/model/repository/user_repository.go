package repository

import (
	"os"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Constantes para o nome da coleção de usuários
const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION" // Nome da coleção de usuários
	DEFAULT_COLLECTION_NAME = "users"                   // Valor padrão hard-coded
)

// getCollectionName retorna o nome da collection com fallback para valor padrão

// Função que retorna o nome da coleção de usuários
func getCollectionName() string { 
	collectionName := os.Getenv(MONGODB_USER_COLLECTION) // Obtém o nome da coleção de usuários do banco de dados
	if collectionName == "" {                            // Se o nome da coleção de usuários não for definido, retorna o nome padrão
		logger.Info("MONGODB_USER_COLLECTION not defined, using default collection name", 
			zap.String("defaultCollection", DEFAULT_COLLECTION_NAME), 
			zap.String("journey", "repository"))                      
		return DEFAULT_COLLECTION_NAME  
	}
	return collectionName 
}

// Função que cria um novo repositório de usuários
func NewUserRepository(
	database *mongo.Database, // Conexão com o banco de dados MongoDB
) UserRepository { // 
	return &userRepository{ // 
		database, // Parametro de conexão com o banco de dados MongoDB
	}
}

// Struct que define o repositório de usuários
type userRepository struct {
	databaseConnection *mongo.Database 
}

// Interface que define os métodos que o repositório de usuários deve implementar
type UserRepository interface {
	CreateUser( // Método que cria um usuário
		userDomain model.UserDomainInterface, 
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser( // Método que atualiza um usuário
		userId string, 
		userDomain model.UserDomainInterface, 
	) *rest_err.RestErr 

	DeleteUser( // Método que deleta um usuário
		userId string, 
	) *rest_err.RestErr 

	FindUserByEmail( // Método que busca um usuário por email
		email string, 
	) (model.UserDomainInterface, *rest_err.RestErr) 

	FindUserByEmailAndPassword( // Método que busca um usuário por email e senha
		email string, 
		password string, 
	) (model.UserDomainInterface, *rest_err.RestErr) 

	FindUserByID( // Método que busca um usuário por ID
		id string, 
	) (model.UserDomainInterface, *rest_err.RestErr) 
}
