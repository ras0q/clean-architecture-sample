SHELL     := /bin/bash
APP_PORT 	:= 8081
APP_INFRA := origin # ginを使うときは`make all APP_INFRA=alt`とする

# Run on Docker
all: stop up logs
stop:
	@docker compose stop
up:
	@APP_INFRA=$(APP_INFRA) docker compose up --build -d
logs:
	@docker compose logs -f

# Run locally
.PHONY: dev
dev:
	@docker compose up db -d --no-recreate
	@APP_PORT=$(APP_PORT) go run github.com/cosmtrek/air@v1.27.3

# Test go-files
.PHONY: test
test:
	@docker compose exec app go test -cover -shuffle=on -v ./...

# Generate templates
gen:
	@go generate ./...
