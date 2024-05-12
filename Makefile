run: build
	@./bin/grt

build:
	@go build -o bin/grt main.go

test:
	@go test -v ./...

mod:
	@go mod tidy


