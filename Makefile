BINARY_DIR=bin

.PHONY: build
build:
	go build -o ${BINARY_DIR}/api cmd/app/main.go

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: swagger
swagger:
	swag init -g cmd/app/main.go