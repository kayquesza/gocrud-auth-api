package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
)

// Variável que define a chave secreta do JWT
var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

// Função que gera um token de autenticação
func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) { // Retorna um token e um erro
	secret := os.Getenv(JWT_SECRET_KEY) // Obtém a chave secreta do JWT

	claims := jwt.MapClaims{ // Uma struct que compõem o que será usado na criação do token
		"id":    ud.ID,                                 // ID do usuário
		"email": ud.email,                              // Email do usuário
		"name":  ud.name,                               // Nome do usuário
		"age":   ud.age,                                // Idade do usuário
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Tempo de expiração do token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Cria um novo token com as claims

	tokenString, err := token.SignedString([]byte(secret)) // Assina o token com a chave secreta
	if err != nil {                                        // Se houver algum erro, retorna um erro
		return "", rest_err.NewInternalServerError( // Retorna um erro interno do servidor
			fmt.Sprintf("Error trying to generate JWT Token. Err=%s", err.Error())) // Retorna um erro interno do servidor
	}

	return tokenString, nil // Retorna o token e nil
}

// Função que verifica se o token é válido
func VerifyTokenMiddleware(c *gin.Context) { // Retorna um erro
	secret := os.Getenv(JWT_SECRET_KEY)                                     // Obtém a chave secreta do JWT
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization")) // Remove o prefixo "Bearer " do token

	token, err := jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) { // Analisa o token
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok { // Verifica se o método de assinatura é HMAC
			return []byte(secret), nil // Retorna a chave secreta
		}

		return nil, rest_err.NewBadRequestError("Invalid token") // Retorna um erro de token inválido
	})

	if err != nil { // Se houver algum erro, retorna um erro
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token") // Retorna um erro de token inválido
		c.JSON(errRest.Code, errRest)                                    // Retorna o erro
		c.Abort()                                                        // Interrompe a requisição
		return                                                           // Retorna
	}

	claims, ok := token.Claims.(jwt.MapClaims) // Obtém as claims do token
	if !ok || !token.Valid {                   // Verifica se o token é válido
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token") // Retorna um erro de token inválido
		c.JSON(errRest.Code, errRest)                                    // Retorna o erro
		c.Abort()                                                        // Interrompe a requisição
		return                                                           // Retorna
	}

	userDomain := userDomain{ // Cria um novo domínio de usuário
		ID:    claims["id"].(string),         // ID do usuário
		email: claims["email"].(string),      // Email do usuário
		name:  claims["name"].(string),       // Nome do usuário
		age:   int8(claims["age"].(float64)), // Idade do usuário
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain)) // Loga o usuário autenticado
	return                                                          // Retorna
}

// Função que remove o prefixo "Bearer " do token
func RemoveBearerPrefix(token string) string { // Retorna o token sem o prefixo "Bearer "
	return strings.TrimPrefix(token, "Bearer ") // Retorna o token sem o prefixo "Bearer "
}
