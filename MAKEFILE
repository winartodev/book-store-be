.PHONY: test build

test: 
	go test -coverprofile=cover.out ./...

build:
	go build -v app/main.go
