start: build
	@ ./bin/main

build:
	@go build -o ./bin ./cmd/main.go


docker-start:
	@docker compose up --build

docker-stop:
	@docker-compose rm -v --force stop
	@docker rmi ticket-booking

