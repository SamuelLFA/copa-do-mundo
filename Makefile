BINARY_DIR=bin

.PHONY: build
build:
	go build -o ${BINARY_DIR}/api cmd/app/main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: deps
deps: tidy vendor

.PHONY: clean
clean:
	rm -rf ${BINARY_DIR}
	go clean --cache -modcache

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: swagger
swagger:
	swag init -g cmd/app/main.go

.PHONY: docker
docker:
	docker-compose -f docker/docker-compose.yaml up -d