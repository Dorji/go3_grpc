package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vlasove/proj/github.com/vlasove/Lec3/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Starting gRPC server...")

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen from  port: %v", err)
	}
	//Our Service
	serv := chat.Server{}
	//Basic gRPC Server
	grpcServer := grpc.NewServer()
	//Ассоциируем gRPC сервер с нашим сервисом
	chat.RegisterChatServiceServer(grpcServer, &serv)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("falied to server with that point: %s", err)
	}

	//go get -u github.com/golang/protobuf/protoc-gen-go
}
