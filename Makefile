build:
	 go build -o ./app ./cmd/service

lint:
	golangci-lint run ./...

docker-build:
	docker build -t imageservice -f ./deploy/docker-compose/Dockerfile ./

run:
	docker compose --file deploy/docker-compose/docker-compose.yaml --project-directory ./  up -d

stop:
	docker compose --file deploy/docker-compose/docker-compose.yaml --project-directory ./  stop