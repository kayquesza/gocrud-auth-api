package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variáveis de ambiente (.env) para o banco de dados MongoDB
var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

// Função para criar uma conexão com o banco de dados MongoDB
func NewMongoDBConnection( // Função que recebe
	ctx context.Context, // o contexto da requisição
) (*mongo.Database, error) { // e retorna uma conexão com o banco de dados ou um erro
	mongodb_uri := os.Getenv(MONGODB_URL)          // Obtém a URI do banco de dados através da v.A (variável de ambiente)
	mongodb_database := os.Getenv(MONGODB_USER_DB) // Obtém o nome do banco de dados através da v.A

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri)) // Tenta se conectar ao banco de dados
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil { // Verifica se a conexão foi estabelecida com sucesso
		return nil, err
	}

	return client.Database(mongodb_database), nil // Retorna a instância do banco específico e nil para erro
}
