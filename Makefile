.PHONY: all vendor rebuild dev logs stop down

all: envInit vendor dev logs

envInit:
	@[[ ! -f .env ]] && cp .env.sample .env || :

vendor:
	@go mod vendor
	@go mod tidy

rebuild: vendor
	@docker-compose up -d --build api

dev: vendor
	@docker-compose up -d

logs:
	# || true silences make error on ctr-c to exit logs
	@docker-compose logs -f api || true

stop:
	@docker-compose stop

down:
	@docker-compose down
