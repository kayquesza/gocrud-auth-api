package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Struct que define a entidade de usuário
type UserEntity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"` // Se o campo for vázio, ignora o campo
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int8               `bson:"age,omitempty"` // omitempty impede que o usuário deixe o campo vázio / mantém o que já tem no banco de dados
}
