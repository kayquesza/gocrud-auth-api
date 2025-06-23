package repository

import (
	"context"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Função que deleta um usuário
func (ur *userRepository) DeleteUser(
	userId string, // ID do usuário
) *rest_err.RestErr { // Retorna um erro
	logger.Info("Initiating deleteUser repository", // Mensagem de log
		zap.String("journey", "deleteUser")) // Jornada da deleção de um usuário

	collection_name := getCollectionName()                          // Obtém o nome da coleção de usuários do banco de dados
	collection := ur.databaseConnection.Collection(collection_name) // Cria uma referência à coleção de usuários do banco de dados

	userIdHex, _ := primitive.ObjectIDFromHex(userId) // Converte o ID do usuário para um ObjectID

	filter := bson.D{{Key: "_id", Value: userIdHex}} // Cria um filtro para deletar o usuário

	_, err := collection.DeleteOne(context.Background(), filter) // Deleta o usuário do banco de dados
	if err != nil {                                              // Se houver algum erro, retorna um erro
		logger.Error("Error trying to update user", err, // Mensagem de log
			zap.String("journey", "deleteUser")) // Jornada da deleção de um usuário
		return rest_err.NewInternalServerError(err.Error()) // Retorna um erro interno do servidor
	}

	logger.Info("deleteUser repository executed successfully", // Mensagem de log
		zap.String("userId", userId),        // ID do usuário
		zap.String("journey", "deleteUser")) // Jornada da deleção de um usuário
	return nil
}
