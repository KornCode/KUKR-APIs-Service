build:
	go build -o server main.go

build.linux:
	GOOS=linux GOARCH=amd64 go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

download:
	go mod download

tidy:
	go mod tidy

d.up:
	docker-compose up

d.down:
	docker-compose down

d.up.build:
	docker-compose up --build
