.PHONY:proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto


.PHONY:build 
build:
	sudo docker build -t client .

.PHONY:run
run:
	sudo docker run -p 50053:50051  -e MICRO_SERVER_ADDRESS=:50051 client

.DEFAULT_GOAL := proto