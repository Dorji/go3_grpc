package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/vlasove/Lec13/userserver/proto/user"
)

type authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

//Контроллер
type handler struct {
	repository   Repository
	tokenService authable
}

func (s *handler) Create(ctx context.Context, in *pb.User, out *pb.Response) error {
	log.Println("user:", in)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//Присланному юзеру перетираем пароль
	in.Password = string(hashedPass)

	if err := s.repository.Create(ctx, MarshallUser(in)); err != nil {
		return err
	}
	//Убираем пароль у пользователя котоырй пойдет в ответ
	in.Password = ""
	out.User = in
	return nil

}

func (s *handler) Get(ctx context.Context, in *pb.User, out *pb.Response) error {
	result, err := s.repository.Get(ctx, in.Id)
	if err != nil {
		return err
	}
	user := UnmarshallUser(result)
	out.User = user
	return nil

}

func (s *handler) GetAll(ctx context.Context, in *pb.Request, out *pb.Response) error {
	results, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	users := UnmarshallUserCollection(results)
	out.Users = users
	return nil
}

func (s *handler) Auth(ctx context.Context, in *pb.User, out *pb.Token) error {
	//Пытаемся узнать есть ли в БД юзер с таким адресом почты
	user, err := s.repository.GetByEmail(ctx, in.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return err
	}
	//Подготавливаем токен
	token, err := s.tokenService.Encode(in)
	if err != nil {
		return err
	}
	out.Token = token
	return nil

}

func (s *handler) ValidateToken(ctx context.Context, in *pb.Token, out *pb.Token) error {
	claims, err := s.tokenService.Decode(in.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == "" {
		return errors.New("invalid user/token pair, try Auth again")
	}

	out.Valid = true
	return nil
}
