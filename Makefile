.PHONY: help coverage mock test build

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

VERSION = $(shell git describe --always --dirty)

## Show Help make command
help:
	@echo ""
	@echo "Usage:"
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
        helpMessage = match(lastLine, /^## (.*)/); \
        if (helpMessage) { \
            helpCommand = substr($$1, 0, index($$1, ":")-1); \
                        helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
                        printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
                } \
        } \
    	{ lastLine = $$0 }' $(MAKEFILE_LIST)
## Generate mock stub
mock:
	mockery --dir ./internal --output internal/mocks --all --keeptree

## Running Unit-test Code
test:
	ENV=test go test `go list ./... | grep -v mocks` -race -coverprofile coverage.cov -cover ./... && go tool cover -func coverage.cov

## Running Code Coverage
coverage:
	ENV=test go test `go list ./... | grep -v mocks` -coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html

## Running Code Dependency Check
check:
	go mod tidy
	go mod download
	go mod verify

## Running code at local to debug propose
run:
	go run main.go

## Build Code to Binary Artifact
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o build/bin main.go
