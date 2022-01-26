.PHONY: test build mod cover coverage start

start: 
	go run ./...

test: 
	go test -coverprofile=cover.out ./...

build:
	go build -v app/main.go

mod: 
	go mod tidy
	go mod download

cover: 
	go tool cover -func=cover.out | grep total

coverage: 
	go tool cover -html=cover.out

migration:
	docker run -d -p 5432:5432 --network bookstore_network --network-alias host --name bookstore_db -e POSTGRES_PASSWORD=postgres -v bookstore_volume:/var/lib/postgresql/data postgres:14
	
docker_build: 
	docker build . -t book-store-api -f deploy/api/Dockerfile

docker_run: 
	docker run -d -p 8080:8080 --network bookstore_network -e DATABASE_HOST=host --name book-store-api book-store-api
