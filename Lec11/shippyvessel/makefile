.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto 

.PHONY:build
build:
	sudo docker build -t shippyvessel .

.PHONY:run
run:
	sudo docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 shippyvessel

.DEFAULT_GOAL := proto