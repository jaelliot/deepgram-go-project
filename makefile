build:
	@go build -o bin/deepgram-go-project

run: build
	@./bin/deepgram-go-project

test:
	@go test -v ./...