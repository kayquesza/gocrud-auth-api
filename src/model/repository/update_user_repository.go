package repository

import (
	"context"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Função que atualiza um usuário
func (ur *userRepository) UpdateUser(
	userId string, // ID do usuári
	userDomain model.UserDomainInterface, 
) *rest_err.RestErr { 
	logger.Info("Initiating updateUser repository", 
		zap.String("journey", "updateUser")) 

	collection_name := getCollectionName()                          // Obtém o nome da coleção de usuários do banco de dados
	collection := ur.databaseConnection.Collection(collection_name) // Cria uma referência à coleção de usuários do banco de dados

	value := converter.ConvertDomainToEntity(userDomain) // Converte o domínio de usuário para uma entidade de usuário
	userIdHex, _ := primitive.ObjectIDFromHex(userId)    // Converte o ID do usuário para um ObjectID

	filter := bson.D{{Key: "_id", Value: userIdHex}} // Cria um filtro para atualizar o usuário
	update := bson.D{{Key: "$set", Value: value}}    // Cria um update para atualizar o usuário

	_, err := collection.UpdateOne(context.Background(), filter, update) // Atualiza o usuário no banco de dados
	if err != nil {                                                      
		logger.Error("Error trying to update user", err, 
			zap.String("journey", "updateUser")) 
		return rest_err.NewInternalServerError(err.Error()) // Retorna um erro interno do servidor
	}

	logger.Info("updateUser repository executed successfully", 
		zap.String("userId", userId),        
		zap.String("journey", "updateUser")) 
	return nil // Retorna nil
}
