.PHONY:proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. proto/dop/dop.proto 

.PHONY:build
build:
	sudo docker build -t dop . 

.PHONY:run
run:
	sudo docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 dop

.DEFAULT_GOAL := proto