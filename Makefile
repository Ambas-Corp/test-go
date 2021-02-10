#Go Params
GOCMD=go
GOBUILD=$(GOCMD) build

CMDDIR=cmd
BIN_NAME=card


all: build run
build:
	rm -f $(BIN_NAME)
	$(GOBUILD) -o $(BIN_NAME) main.go
proto:
	protoc --go_out=plugins=grpc:. handler/grpc/proto/card.proto
run:
	./$(BIN_NAME)
