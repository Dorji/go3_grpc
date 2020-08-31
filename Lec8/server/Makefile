.PHONY:proto
proto:
	protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto


.PHONY:build 
build:
	sudo docker build -t server .

.PHONY:run
run:
	sudo docker run -p 50051:50051 server 

.DEFAULT_GOAL := proto