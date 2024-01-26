package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	permissions["user_id"] = id                               //Define a validade do token
	// Cria o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	// Assina o token
	return token.SignedString([]byte(config.SecretKey))
	//return token.SignedString(config.SecretKey)
}

func ValidadeToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}
	//Retorna todas as permissões e verifica se está tudo certo e token ainda é valido
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	//Verifica se o token possui mais de uma palavra e pega sempre a segunda. EX: Bearer rfeojqjfri
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	//verifica se o metodo de assinatura utilizado pertence a familia correta
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %s", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Converte para o formato do JWT, depois para um string e depois para um uint
		ID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return ID, nil
	}
	return 0, errors.New("token invalido")
}
