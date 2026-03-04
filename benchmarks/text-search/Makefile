.PHONY: up down clean logs ps install-extensions

up:
	docker compose up -d
	@echo "Waiting for services to be healthy..."
	@sleep 15

down:
	docker compose down

clean:
	docker compose down -v

logs:
	docker compose logs -f

ps:
	docker compose ps

install-extensions:
	docker exec -i benchmark-postgres psql -U benchmark -d benchmark < scripts/install_extensions.sql
