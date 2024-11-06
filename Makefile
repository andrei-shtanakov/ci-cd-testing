.PHONY: build test run docker

build:
	go build -o app ./cmd/app

test:
	go test ./...

run: build
	./app

docker:
	docker build -t myapp .