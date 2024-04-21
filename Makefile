run: build
	@./bin/api

build:
	@go build -o bin/api ./src/

test:
	@go test -v ./...