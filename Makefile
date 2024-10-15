build:
	@go build -o bin/tikservice

run: build
	@./bin/tikservice