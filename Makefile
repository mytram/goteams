OUT = out
SHORT_SHA ?= $(shell git rev-parse --short HEAD)
SEMVER = $(shell git describe --tags --abbrev=0)
TAG = $(or $(SEMVER),$(SHORT_SHA))
DSN ?= "host=localhost user=postgres password=postgres dbname=goteams port=5432 sslmode=disable TimeZone=UTC"

GO_LDFLAS = -X main.Version=$(TAG)

$(OUT):
	@mkdir -p $@

build: | $(OUT)
	$(GO_ENV) go build -o $(OUT) -ldflags='$(GO_LDFLAGS)' ./cmd/...

migrate: build
	@DSN=$(DSN) $(OUT)/migrate

server: build
	@DSN=$(DSN) $(OUT)/server

.PHONY: lint
lint:
	golangci-lint run

.PHONY: cleanup-db
cleanup-db:
	docker rm -f $(shell docker ps -q --filter publish=5432) > /dev/null 2>&1 || true

.PHONY: docker-run-postgres
docker-run-postgres: cleanup-db ## Spin up a local postgres instance, for local development outside of docker-compose
	docker run -d --rm --name postgres \
		-p 5432:5432 \
		-e POSTGRES_DB=goteams \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_USER=postgres \
		postgres:12-alpine
