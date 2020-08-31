package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/Lec9/server/proto/consignment"
)

// const (
// 	port = ":50051"
// )

type repository interface {
	Create(*pb.Command) (*pb.Command, error)
	GetAll() []*pb.Command
}

//Repository ... Наша база данных
type Repository struct {
	commands []*pb.Command
}

//Create ....
func (r *Repository) Create(command *pb.Command) (*pb.Command, error) {
	updatedCommands := append(r.commands, command)
	r.commands = updatedCommands
	return command, nil
}

//GetAll ...
func (r *Repository) GetAll() []*pb.Command {
	return r.commands
}

type service struct {
	repo repository
}

func (s *service) CreateCommand(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	command, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Request to create response: %v", command)
	return &pb.Response{Created: true, Command: command}, nil
}

//GetAllCommands ...
func (s *service) GetAllCommands(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	commands := s.repo.GetAll()
	return &pb.Response{Commands: commands}, nil
}

func main() {
	repo := &Repository{}

	srvService := micro.NewService(
		micro.Name("server"),
	)
	srvService.Init()

	ourService := &service{repo}
	if err := pb.RegisterShippingServiceHandler(srvService.Server(), ourService); err != nil {
		log.Panic(err)
	}

	if err := srvService.Run(); err != nil {
		log.Panic(err)
	}

	// //Настройка gRPC сервера
	// listener, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen port: %v", err)
	// }

	// server := grpc.NewServer()

	// //Регистрируем наш сервис для сервера
	// ourService := &service{repo}
	// pb.RegisterShippingServiceServer(server, ourService)
	// //Чтобы выходные параметры сервера сохранялись в go-runtime
	// reflection.Register(server)

	// log.Println("gRPC server running on port:", port)
	// if err := server.Serve(listener); err != nil {
	// 	log.Fatalf("failed to server from port: %v", err)
	// }
}
