PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

GOBUILD=CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

WORKSPACE ?= name

build:
	cd ./cmd/$(WORKSPACE) && $(GOBUILD)

# swagger
swagger:
	swag init --pd -d ./cmd/server -o ./cmd/server/docs

# openapi 3.0
openapi: swagger
	cd ./cmd/server && klctl openapi  -f ./docs/swagger.json

# client
client:
	klctl gen client server  --output ./cmd/client --spec-url http://127.0.0.1/example-server