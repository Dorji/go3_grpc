.PHONY:cli 
cli:
	evans -r repl -p 50051


.PHONY:proto
proto:
	protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto


.PHONY:build 
build:
	sudo docker build -t client .

.PHONY:run
run:
	sudo docker run -p 50052:50051 client 

.DEFAULT_GOAL := proto