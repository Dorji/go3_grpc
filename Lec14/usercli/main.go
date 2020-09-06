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
	rsp, err := client.Create(ctx, user)
	//respUsers, err := client.GetAll(ctx, &proto.Request{})
	if err != nil {
		return err
	}

	fmt.Println("Response: ", rsp.User)
	//fmt.Println("All users in database:", respUsers.Users)

	return nil
}

func main() {

	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your Name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "E-Mail",
			},
			&cli.StringFlag{
				Name:  "company",
				Usage: "Company Name",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Password",
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
				log.Println("error creating user: ", err.Error())
				return err
			}

			return nil
		}),
	)
}
