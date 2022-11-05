server:
	go run src/cmd/main.go

build:
	go build -o bin/server src/cmd/main.go

d.up:
	docker-compose up

d.down:
	docker-compose down

d.up.build:
	docker-compose --build up