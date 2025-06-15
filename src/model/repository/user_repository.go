package repository

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

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

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
