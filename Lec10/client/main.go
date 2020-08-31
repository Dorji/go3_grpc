package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/Lec10/client/proto/consignment"
)

const (
	//address         = "localhost:50051"
	defaultFilename = "command.json"
)

func parseJSON(file string) (*pb.Command, error) {
	var command *pb.Command
	fileBody, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(fileBody, &command)
	return command, err
}

func main() {

	// connection, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("can not connect to port: %v", err)
	// }
	// defer connection.Close()

	// client := pb.NewShippingServiceClient(connection)
	service := micro.NewService(
		micro.Name("client"),
	)
	service.Init()
	//Сопрягаем наш текущий сервис в качесвте клиента с сервисом по имени "server"
	//внутри MDNS
	client := pb.NewShippingServiceClient("server", service.Client())

	command, err := parseJSON(defaultFilename)
	if err != nil {
		log.Fatalf("can not parse .json file: %v", err)
	}
	resp, err := client.CreateCommand(context.Background(), command)
	if err != nil {
		log.Fatalf("can not get response: %v", err)
	}
	log.Printf("Created: %t", resp.Created)
	log.Printf("Body: %v", resp.Command)

	getAll, err := client.GetAllCommands(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("can not get response: %v", err)
	}
	for _, v := range getAll.Commands {
		log.Println(v)
	}

}
