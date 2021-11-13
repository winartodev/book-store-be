.PHONY: test build create_db remove_db stop_db start_db

test: 
	go test -coverprofile=cover.out ./...

build:
	go build -v app/main.go

#  thos command for control the databases
create_db: 
	cd db; 	echo "Create Database..."; \
	docker-compose up -d

remove_db:
	cd db; echo "Database Down..."; \
	docker-compose down

stop_db:
	cd db; echo "Database Stopped"; \
	docker-compose stop

start_db: 
	cd db; echo "Database Started"; \
	docker-compose start 