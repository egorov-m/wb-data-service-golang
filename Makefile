.PHONY: help
help:
		@echo "USAGE"
		@echo "    make <commands>"
		@echo ""
		@echo "AVAILABLE COMMANDS"
		@echo "up          Create and start containers in the background."
		@echo "down        Stop and remove containers, networks."
		@echo "swag-gen    Generation of Swagger documentation for the service."
		@echo "run-service Running wb service."
		@echo "run-worker  Running wb worker."

.PHONY: up
up:
		docker-compose -f docker-compose.yaml up -d

.PHONY: down
down:
		docker-compose -f docker-compose.yaml down && docker network prune --force

.PHONY: swag-gen
swag-gen:
		cd wb-data-service && go run github.com/swaggo/swag/cmd/swag init -g ./cmd/wb-data-service/main.go -o ./docs && cd ..

.PHONY: run-service
run-service:
		cd wb-data-service && go run ./cmd/wb-data-service/main.go && cd ..

.PHONY: run-worker
run-worker:
		cd wb-data-worker && go run ./cmd/wb-data-worker/main.go && cd ..
