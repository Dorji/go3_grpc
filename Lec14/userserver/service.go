package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/vlasove/Lec14/userserver/proto/user"
)

var (
	//Ключ подписи - ключ/метка
	//НИКОГДА НЕ ИСПОЛЬЗУЙТЕ СОБСТВЕННЫЕ КЛЮЧИ.
	key = []byte("mySuperSecretKey")
)

// Токен-объект == метаданные, эти данные будут хешироваться и участвовать в
//генерации токена
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

//Превращает токен в метаданные
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//Конструируем метаданные и возвращаем токен
func (srv *TokenService) Encode(user *pb.User) (string, error) {

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "userserver",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
