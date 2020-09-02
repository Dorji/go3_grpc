package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	micro "github.com/micro/go-micro"
	pb "github.com/vlasove/Lec12/shippyclient/proto/consignment"
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
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		nicePrint(v)
	}
}
