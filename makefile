.PHONY: default run build docs clean

APP_NAME=gin-boilerplate

default: run

run:
	@swag init -g cmd/main.go
	@go run cmd/main.go
build:
	@go build -o $(APP_NAME) cmd/main.go
docs:
	@swag init -g cmd/main.go
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs