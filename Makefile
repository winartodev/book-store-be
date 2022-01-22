.PHONY: test build mod create_db remove_db stop_db start_db cover coverage run

run: 
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
	
#  this command for control the databases
create_db: 
	cd db; 	echo "Create Database..."; \
	docker-compose up -d

remove_db:
	cd db; echo "Database Down..."; \
	docker-compose down -v

stop_db:
	cd db; echo "Database Stopped"; \
	docker-compose stop

start_db: 
	cd db; echo "Database Started"; \
	docker-compose start 

# DOCKER
docker_build: 
	docker build . -t book-store-be -f deploy/api/Dockerfile

docker_run: 
	docker run --name book-store-api -dp 8081:8080 book-store-be

docker_log:
	docker logs -f book-store-api

docker_remove_container: 
	docker rm -f book-store

docker_force_all:
	docker rm -f book-store; \
	docker rmi book-store-api