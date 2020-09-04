package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/vlasove/Lec13/userserver/proto/user"
)

// message User {
//     string id = 1;
//     string name = 2;
//     string email = 3;
//     string company = 4;
//     string password = 5;
// }

const TableScheme = `
	create table if not exists users (
		id varchar(30) primary key not null,
		name varchar(125) not null,
		email varchar(125) not null,
		password varchar(256) not null,
		company varchar(125)
	);
`

func main() {
	db, err := NewConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
		log.Panic(err)
	}
	defer db.Close()

	db.MustExec(TableScheme)

	repo := NewPostgresRepository(db)

	tokenService := &TokenService{repo}

	service := micro.NewService(
		micro.Name("userserver"),
		micro.Version("latest"),
	)

	service.Init()

	controller := &handler{repo, tokenService}
	if err := pb.RegisterUserServiceHandler(service.Server(), controller); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}
