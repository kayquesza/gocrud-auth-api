package repository

import (
	"context"
	"os"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initiating CreateUser method in UserRepository",
		zap.String("journey", "createUser"))

	colletion_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(colletion_name)

	value := converter.ConvertDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error inserting user into MongoDB", err,
			zap.String("journey", "createUser"))
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("Creating user in MongoDB",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"))
	return converter.ConvertEntityToDomain(*value), nil

}
