package repository

import (
	"context"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Função que cria um usuário
func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface, // Domínio de usuário
) (model.UserDomainInterface, *rest_err.RestErr) { // Retorna o domínio de usuário e um erro
	logger.Info("Initiating CreateUser method in UserRepository", // Mensagem de log
		zap.String("journey", "createUser")) // Jornada da criação de um usuário

	collection_name := getCollectionName()                          // Obtém o nome da coleção de usuários do banco de dados
	collection := ur.databaseConnection.Collection(collection_name) // Cria uma referência à coleção de usuários do banco de dados

	value := converter.ConvertDomainToEntity(userDomain)             // Converte o domínio de usuário para uma entidade de usuário
	result, err := collection.InsertOne(context.Background(), value) // Insere o usuário no banco de dados
	if err != nil {                                                  // Se houver algum erro, retorna um erro
		logger.Error("Error inserting user into MongoDB", err, // Mensagem de log
			zap.String("journey", "createUser")) // Jornada da criação de um usuário
	}

	value.ID = result.InsertedID.(primitive.ObjectID) // Obtém o ID do usuário inserido

	logger.Info("Creating user in MongoDB", // Mensagem de log
		zap.String("userId", value.ID.Hex()), // ID do usuário
		zap.String("journey", "createUser"))  // Jornada da criação de um usuário
	return converter.ConvertEntityToDomain(*value), nil // Converte a entidade de usuário para um domínio de usuário

}
