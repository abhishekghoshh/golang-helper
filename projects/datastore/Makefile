.SILENT:

.PHONY: build test image

IMAGE_NAME = abhishek1009/datastore
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
	docker build -t $(IMAGE_NAME):$(DEFAULT_TAG) -f Dockerfile .

push_image: image
	docker image push $(IMAGE_NAME):$(DEFAULT_TAG)

image_arm:
	docker build --build-arg TARGETARCH=arm64 --build-arg TARGETOS=linux -t $(IMAGE_NAME):arm-$(DEFAULT_TAG) -f deployment/Dockerfile-advanced .

push_image_arm: image_arm
	docker image push $(IMAGE_NAME):arm-$(DEFAULT_TAG)

image_amd:
	docker build --build-arg TARGETARCH=amd64 --build-arg TARGETOS=linux -t $(IMAGE_NAME):amd-$(DEFAULT_TAG) -f deployment/Dockerfile-advanced .

push_image_amd: image_amd
	docker image push $(IMAGE_NAME):amd-$(DEFAULT_TAG)

air:
	air -c .air.toml