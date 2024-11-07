.SILENT:

.PHONY: build test image

IMAGE_NAME = abhishek1009/group-chat
DEFAULT_TAG = latest


say_hello:
	cat ./Makefile

build:
	go build cmd/chat/main.go

run:
	go run cmd/chat/main.go

test:
	go test ./...

image:
	docker build -t $(IMAGE_NAME):$(DEFAULT_TAG) -f deployment/Dockerfile .
air:
	air -c .air.toml