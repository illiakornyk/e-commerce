build:
	@go build -o bin/e-commerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/e-commerce
