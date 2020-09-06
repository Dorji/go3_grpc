package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	pb "github.com/vlasove/Lec15/shippyclient/proto/consignment"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

// nicePrint ...
func nicePrint(c *pb.Consignment) {
	log.Println("==================CONSIGNMENT DATA===============")
	log.Printf("Description: %s\n", c.GetDescription())
	log.Printf("Weight: %d\n", c.GetWeight())
	log.Printf("Containers: %v\n", c.GetContainers())
	log.Printf("VesselID: %s\n", c.GetVesselId())
	log.Println("=================================================")
	log.Println()
}

func main() {

	service := micro.NewService(micro.Name("shippyclient"))
	service.Init()

	client := pb.NewShippingServiceClient("shippyserver", service.Client())

	file := defaultFilename
	//Здесь начинаем добавлять аутентификацию
	//Пусть client получает токен из CLI
	var token string
	log.Println(os.Args) // пробросит в терминал все аргументы CLI

	if len(os.Args) < 3 {
		log.Fatal("not enough args. Expected file_path and token")
	}
	//os.Args[0] == "main.go"
	file = os.Args[1]
	token = os.Args[2]

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	//Создаем новй контекст который будет содержать токен
	//После чего мы будем передавать этот контекст серверу
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	//Пробуем осуществить create запрос для shippyserver

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		nicePrint(v)
	}
}
