build:
	@go build -o ./bin/http

run: build
	@./bin/http

test:
	@go test ./...
