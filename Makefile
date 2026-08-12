.PHONY: start lint

start:
	go run main.go

lint:
	golangci-lint run
