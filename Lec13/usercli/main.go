package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	proto "github.com/vlasove/Lec13/usercli/proto/user"
)

func createUser(ctx context.Context, service micro.Service, user *proto.User) error {
	client := proto.NewUserService("userserver", service.Client())
	response, err := client.Create(ctx, user)
	if err != nil {
		return err
	}
	fmt.Println("Response:", response.User)
	return nil
}

func main() {
	service := micro.NewService(
		//Информацию о пользователе забираем из командной строки в качестве cli параметров при запуске контейнера
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			&cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			log.Println(c)
			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			log.Println("test:", name, email, company, password)
			ctx := context.Background()
			user := &proto.User{
				Name:     name,
				Email:    email,
				Company:  company,
				Password: password,
			}
			if err := createUser(ctx, service, user); err != nil {
				log.Println("error while user creation:", err.Error())
				return err
			}
			return nil
		}),
	)
}
