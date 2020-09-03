.PHONY:proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. proto/vessel/vessel.proto

.PHONY:build
build:
	sudo docker build -t shippyvessel .

.PHONY:run
run:
	sudo docker run -p 50054:50051 -e MICRO_SERVER_ADDRESS=:50051 shippyvessel

.DEFAULT_GOAL := proto