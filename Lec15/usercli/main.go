package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	microclient "github.com/micro/go-micro/v2/client"
	pb "github.com/vlasove/Lec14/usercli/proto/user"
)

func main() {

	srv := micro.NewService(

		micro.Name("usercli"),
		micro.Version("latest"),
	)

	srv.Init()

	//создаем функционал клиента
	client := pb.NewUserService("userserver", microclient.DefaultClient)

	////Кого будем скармливать БД
	name := "Evgen Petrov"
	email := "evgeny_vlasov@gmail.com"
	password := "123456789q"
	company := "Evgen.IO"

	user := &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	}
	//Доабвляем юзера в БД
	resp, err := client.Create(context.TODO(), user)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}

	log.Printf("User created: %s", resp.User.Name)

	allUsers, err := client.GetAll(context.TODO(), &pb.Request{})
	if err != nil {
		log.Fatalf("DB empty/could not listing all users")
	}

	log.Printf("Find %v users in DB", allUsers)
	//RPC вызов аутентификации (пытается найти юзера по емейл в БД)
	securityUser := &pb.User{
		Email:    email,
		Password: password,
	}
	authResponse, err := client.Auth(context.TODO(), securityUser)
	if err != nil {
		log.Fatalf("can not auth for user: %v", err)
	}
	//В случае если все ок - выбитый токен кидаем в log
	log.Printf("Access token: %s     \n", authResponse.Token)
	return

}
