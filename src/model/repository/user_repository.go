package repository

import (
	"os"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
	DEFAULT_COLLECTION_NAME = "users" // Valor padrão hard-coded
)

// getCollectionName retorna o nome da collection com fallback para valor padrão
func getCollectionName() string {
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	if collectionName == "" {
		logger.Info("MONGODB_USER_COLLECTION not defined, using default collection name",
			zap.String("defaultCollection", DEFAULT_COLLECTION_NAME),
			zap.String("journey", "repository"))
		return DEFAULT_COLLECTION_NAME
	}
	return collectionName
}

func NewUserRepository(
	database *mongo.Database, // Conexão com o banco de dados MongoDB
) UserRepository {
	return &userRepository{
		database, // Parametro de conexão com o banco de dados MongoDB
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	DeleteUser(
		userId string,
	) *rest_err.RestErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPassword(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
