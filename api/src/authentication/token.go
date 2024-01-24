package authentication

import (
	"time"

	"github.com/caio1459/devbook/src/config"
	jwt "github.com/golang-jwt/jwt/v5"
)

// Retorna o token com as permissões do usuário
func GenerateToken(id uint64) (string, error) {
	// Gera permições
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true                          //Usuário autorizado
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() //Validade de 6 horas do token
	permissions["userId"] = id                                //Define a validade do token

	// Cria o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	// Assina o token
	signedToken, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
