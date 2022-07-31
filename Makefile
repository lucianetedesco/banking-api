APP=banking-api

## help: show this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## clean: cleans the binary
clean:
	go clean

## test: run unit tests
test:
	go test ./... -v

## build: build the application
build:
	go build -o ${APP} main.go

## run: run the application
run:
	go run main.go