# Makefile for SFTP Indexer app with Docker Compose

APP_NAME=sftp-indexer
IMAGE_NAME=sftp-indexer
COMPOSE_FILE=docker-compose.yml

# Install mockgen
tools: $(TOOLS_BIN)
	go install go.uber.org/mock/mockgen@latest

# Generate mocks
mocks: tools
	mockgen -source=sftpclient/sftpclient.go -destination=mocks/sftp_client.go -package=mocks
	mockgen -source=indexer/client.go -destination=mocks/index_client.go -package=mocks

test: mocks
	go test ./... -v

# Build the Go app image
build:
	docker build -t $(IMAGE_NAME) .

# Start all services
up:
	docker-compose -f $(COMPOSE_FILE) up --build -d

# Stop all services
down:
	docker-compose -f $(COMPOSE_FILE) down

# View logs for all services
logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

# Tail logs for Go app only
logs-app:
	docker-compose -f $(COMPOSE_FILE) logs -f $(APP_NAME)

# Rebuild and restart only the Go app container
restart-app:
	docker-compose -f $(COMPOSE_FILE) up -d --no-deps --build $(APP_NAME)

# Remove built Docker image
clean:
	docker rmi -f $(IMAGE_NAME) || true

indexer:
	docker exec -it sftp-indexer /bin/sh