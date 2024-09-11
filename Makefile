build:
	@go build -o bin/ffpoker
run: build
	@./bin/ffpoker
test:
	go test -v ./...