.PHONY:proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto


.PHONY:build 
build:
	sudo docker build -t server .

.PHONY:run
run:
	sudo docker run -p 50051:50051  -e MICRO_SERVER_ADDRESS=:50051 server 

.DEFAULT_GOAL := proto