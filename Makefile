.PHONY: build

mod:
	go mod tidy
	go mod verify

build:
	go build -ldflags="-s -w"    

docker_vault:
	docker compose -f ./docker-compose.yml up -d vault
