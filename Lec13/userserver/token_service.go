package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/vlasove/Lec13/userserver/proto/user"
)

var (
	key = []byte("MySuperSecretKeyForThisApplication")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

//Middleware для преобразований пользователей
type TokenService struct {
	repo Repository
}

//Декодировать токен (строку) в токен-объект
func (s *TokenService) Decode(token string) (*CustomClaims, error) {
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	//Проверяем валидность токена и возвращаем токен-объект
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//Преобразование токен-объекта в JWT
func (s *TokenService) Encode(user *pb.User) (string, error) {
	//Создаем токен-объект
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 10000,
			Issuer:    "userserver",
		},
	}
	//Подготовка параметров и генерация
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Преобразование в строку и возврат
	return token.SignedString(key)
}
