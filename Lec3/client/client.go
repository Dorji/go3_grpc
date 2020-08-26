package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vlasove/proj/github.com/vlasove/Lec3/chat"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client starting ....")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect to:%v", err)
	}
	defer conn.Close()

	cli := chat.NewChatServiceClient(conn)
	//Дергаем SayHello у сервера
	resp, err := cli.SayHello(context.Background(), &chat.Message{Body: "Hello from Client!"})
	if err != nil {
		log.Fatalf("Error when SayHello: %s", err)
	}
	log.Printf("Response server: %v", resp.Body)
	//Дергаем SayGoodbye у сервера
	resp, err = cli.SayGoodBye(context.Background(), &chat.Message{Body: "Goodbye from Client!"})
	if err != nil {
		log.Fatalf("Error when SayGoodBye: %s", err)
	}
	log.Printf("Response server: %v", resp.Body)

}
