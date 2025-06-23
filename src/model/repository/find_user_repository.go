package repository

import (
	"context"
	"fmt"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository/entity"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Função que busca um usuário por email
func (ur *userRepository) FindUserByEmail(
	email string, // Email do usuário
) (model.UserDomainInterface, *rest_err.RestErr) { // Retorna o domínio de usuário e um erro
	logger.Info("Initiating findUserByEmail method in UserRepository", // Mensagem de log
		zap.String("journey", "findUserByEmail")) // Jornada da busca de um usuário por email

	collection_name := getCollectionName()                          // Obtém o nome da coleção de usuários do banco de dados
	collection := ur.databaseConnection.Collection(collection_name) // Cria uma referência à coleção de usuários do banco de dados

	userEntity := &entity.UserEntity{} // Cria uma entidade de usuário

	filter := bson.D{{Key: "email", Value: email}} // Cria um filtro para buscar o usuário por email
	err := collection.FindOne(
		context.Background(), // Contexto da requisição
		filter,               // Filtro para buscar o usuário por email
	).Decode(userEntity) // Decodifica o usuário encontrado

	if err != nil { // Se houver algum erro, retorna um erro
		if err == mongo.ErrNoDocuments { // Se o usuário não for encontrado, retorna um erro
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email) // Mensagem de erro
			logger.Error(errorMessage, err, // Mensagem de log
				zap.String("journey", "findUserByEmail")) // Jornada da busca de um usuário por email

			return nil, rest_err.NewNotFoundError(errorMessage) // Retorna um erro de usuário não encontrado
		}

		errorMessage := "Error trying to find user by email" // Mensagem de erro
		logger.Error(errorMessage, err,                      // Mensagem de log
			zap.String("journey", "findUserByEmail")) // Jornada da busca de um usuário por email

		return nil, rest_err.NewInternalServerError(errorMessage) // Retorna um erro interno do servidor
	}

	logger.Info("FindUserByEmail repository executed successfully", // Mensagem de log
		zap.String("journey", "findUserByEmail"),  // Jornada da busca de um usuário por email
		zap.String("email", email),                // Email do usuário
		zap.String("userId", userEntity.ID.Hex())) // ID do usuário
	return converter.ConvertEntityToDomain(*userEntity), nil
}

// Função que busca um usuário por ID
func (ur *userRepository) FindUserByID(
	id string, // ID do usuário
) (model.UserDomainInterface, *rest_err.RestErr) { // Retorna o domínio de usuário e um erro
	logger.Info("Initiating findUserByID method in UserRepository", // Mensagem de log
		zap.String("journey", "findUserByID")) // Jornada da busca de um usuário por ID

	collection_name := getCollectionName()                          // Obtém o nome da coleção de usuários do banco de dados
	collection := ur.databaseConnection.Collection(collection_name) // Cria uma referência à coleção de usuários do banco de dados

	userEntity := &entity.UserEntity{} // Cria uma entidade de usuário

	objectId, _ := primitive.ObjectIDFromHex(id)    // Converte o ID do usuário para um ObjectID
	filter := bson.D{{Key: "_id", Value: objectId}} // Cria um filtro para buscar o usuário por ID
	err := collection.FindOne(
		context.Background(), // Contexto da requisição
		filter,               // Filtro para buscar o usuário por ID
	).Decode(userEntity) // Decodifica o usuário encontrado

	if err != nil { // Se houver algum erro, retorna um erro
		if err == mongo.ErrNoDocuments { // Se o usuário não for encontrado, retorna um erro
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", id) // Mensagem de erro
			logger.Error(errorMessage, err, // Mensagem de log
				zap.String("journey", "findUserByID")) // Jornada da busca de um usuário por ID

			return nil, rest_err.NewNotFoundError(errorMessage) // Retorna um erro de usuário não encontrado
		}

		errorMessage := "Error trying to find user by email" // Mensagem de erro
		logger.Error(errorMessage, err,                      // Mensagem de log
			zap.String("journey", "findUserByID")) // Jornada da busca de um usuário por ID

		return nil, rest_err.NewInternalServerError(errorMessage) // Retorna um erro interno do servidor
	}

	logger.Info("findUserByID repository executed successfully", // Mensagem de log
		zap.String("journey", "findUserByID"),     // Jornada da busca de um usuário por ID
		zap.String("userId", userEntity.ID.Hex())) // ID do usuário
	return converter.ConvertEntityToDomain(*userEntity), nil // Converte a entidade de usuário para um domínio de usuário
}

// Função que busca um usuário por email e senha
func (ur *userRepository) FindUserByEmailAndPassword(
	email string, // Email do usuário
	password string, // Senha do usuário
) (model.UserDomainInterface, *rest_err.RestErr) { // Retorna o domínio de usuário e um erro
	logger.Info("Initiating findUserByEmailAndPassword method in UserRepository", // Mensagem de log
		zap.String("journey", "findUserByEmailAndPassword")) // Jornada da busca de um usuário por email e senha

	collection_name := getCollectionName()                          // Obtém o nome da coleção de usuários do banco de dados
	collection := ur.databaseConnection.Collection(collection_name) // Cria uma referência à coleção de usuários do banco de dados

	userEntity := &entity.UserEntity{} // Cria uma entidade de usuário

	filter := bson.D{ // Cria um filtro para buscar o usuário por email e senha
		{Key: "email", Value: email},       // Filtro para buscar o usuário por email
		{Key: "password", Value: password}, // Filtro para buscar o usuário por senha
	}
	err := collection.FindOne( // Busca o usuário por email e senha
		context.Background(), // Contexto da requisição
		filter,               // Filtro para buscar o usuário por email e senha
	).Decode(userEntity) // Decodifica o usuário encontrado

	if err != nil { // Se houver algum erro, retorna um erro
		if err == mongo.ErrNoDocuments { // Se o usuário não for encontrado, retorna um erro
			errorMessage := ("User or password is invalid") // Mensagem de erro
			logger.Error(errorMessage, err,                 // Mensagem de log
				zap.String("journey", "findUserByEmailAndPassword")) // Jornada da busca de um usuário por email e senha

			return nil, rest_err.NewForbiddenError(errorMessage) // Retorna um erro de usuário não encontrado
		}

		errorMessage := "Error trying to find user by email and password" // Mensagem de erro
		logger.Error(errorMessage, err,                                   // Mensagem de log
			zap.String("journey", "findUserByEmailAndPassword")) // Jornada da busca de um usuário por email e senha

		return nil, rest_err.NewInternalServerError(errorMessage) // Retorna um erro interno do servidor
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully", // Mensagem de log
		zap.String("journey", "findUserByEmailAndPassword"), // Jornada da busca de um usuário por email e senha
		zap.String("email", email),                          // Email do usuário
		zap.String("userId", userEntity.ID.Hex()))           // ID do usuário
	return converter.ConvertEntityToDomain(*userEntity), nil // Converte a entidade de usuário para um domínio de usuário
}
