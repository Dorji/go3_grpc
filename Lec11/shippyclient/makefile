.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consigment.proto 

.PHONY:build
build:
	sudo docker build -t shippyclient .

.PHONY:run
run:
	sudo docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 shippyclient

.DEFAULT_GOAL := proto