.PHONY: help
help:
		@echo "USAGE"
		@echo "    make <commands>"
		@echo ""
		@echo "AVAILABLE COMMANDS"
		@echo "up         Create and start containers in the background."
		@echo "down       Stop and remove containers, networks."

.PHONY: up
up:
		docker-compose -f docker-compose.yaml up -d

.PHONY: down
down:
		docker-compose -f docker-compose.yaml down && docker network prune --force
