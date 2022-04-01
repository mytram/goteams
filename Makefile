OUT = out
SHORT_SHA ?= $(shell git rev-parse --short HEAD)
SEMVER = $(shell git describe --tags --abbrev=0)
TAG = $(or $(SEMVER),$(SHORT_SHA))
DB_FILE ?= "db/goteams.db"

GO_LDFLAS = -X main.Version=$(TAG)

$(OUT):
	@mkdir -p $@

build: | $(OUT)
	$(GO_ENV) go build -o $(OUT) -ldflags='$(GO_LDFLAGS)' ./cmd/...

migrate: build
	DB_FILE=$(DB_FILE) $(OUT)/migrate

server: build
	DB_FILE=$(DB_FILE) $(OUT)/server

.PHONY: lint
lint:
	golangci-lint run
