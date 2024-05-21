build:
	@go build -o bin/notes cmd/main.go

test:
	@go test -b ./..

run: build
	@./bin/notes