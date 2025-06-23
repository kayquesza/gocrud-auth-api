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
func getCollectionName() string { // Retorna o nome da coleção de usuários
	collectionName := os.Getenv(MONGODB_USER_COLLECTION) // Obtém o nome da coleção de usuários do banco de dados
	if collectionName == "" {                            // Se o nome da coleção de usuários não for definido, retorna o nome padrão
		logger.Info("MONGODB_USER_COLLECTION not defined, using default collection name", // Mensagem de log
			zap.String("defaultCollection", DEFAULT_COLLECTION_NAME), // Nome da coleção de usuários padrão
			zap.String("journey", "repository"))                      // Jornada da coleção de usuários
		return DEFAULT_COLLECTION_NAME // Retorna o nome da coleção de usuários padrão
	}
	return collectionName // Retorna o nome da coleção de usuários
}

// Função que cria um novo repositório de usuários
func NewUserRepository(
	database *mongo.Database, // Conexão com o banco de dados MongoDB
) UserRepository { // Retorna um repositório de usuários
	return &userRepository{ // Retorna um repositório de usuários
		database, // Parametro de conexão com o banco de dados MongoDB
	}
}

// Struct que define o repositório de usuários
type userRepository struct {
	databaseConnection *mongo.Database // Conexão com o banco de dados MongoDB
}

// Interface que define os métodos que o repositório de usuários deve implementar
type UserRepository interface {
	CreateUser( // Método que cria um usuário
		userDomain model.UserDomainInterface, // Domínio de usuário
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser( // Método que atualiza um usuário
		userId string, // ID do usuário
		userDomain model.UserDomainInterface, // Domínio de usuário
	) *rest_err.RestErr // Retorna um erro

	DeleteUser( // Método que deleta um usuário
		userId string, // ID do usuário
	) *rest_err.RestErr // Retorna um erro

	FindUserByEmail( // Método que busca um usuário por email
		email string, // Email do usuário
	) (model.UserDomainInterface, *rest_err.RestErr) // Retorna o domínio de usuário e um erro

	FindUserByEmailAndPassword( // Método que busca um usuário por email e senha
		email string, // Email do usuário
		password string, // Senha do usuário
	) (model.UserDomainInterface, *rest_err.RestErr) // Retorna o domínio de usuário e um erro

	FindUserByID( // Método que busca um usuário por ID
		id string, // ID do usuário
	) (model.UserDomainInterface, *rest_err.RestErr) // Retorna o domínio de usuário e um erro
}
