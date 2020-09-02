.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consigment.proto 

.PHONY:build
build:
	sudo docker build -t shippyservice .

.PHONY:run
run:
	sudo docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippyservice

.DEFAULT_GOAL := proto