APP_NAME := docker-cli-plugin
VERSION := 0.0.1

build: test
	go build -ldflags="-X 'main.version=${VERSION}'" -o ${APP_NAME} ./cmd/

test:
	go vet ./...
	go test -v ./...
