include .env
export $(shell sed 's/=.*//' .env)

setup-dev:
	go install golang.org/x/tools/gopls@latest && \
    go install github.com/a-h/templ/cmd/templ@latest && \
	go install github.com/air-verse/air@latest && \
	go install github.com/pressly/goose/v3/cmd/goose@latest

bun-build:
	bun run build

bun-watch:
	bun run watch

templ-build:
	templ generate

air-server-build:
	go build -o ./tmp/app cmd/app/main.go

dev-server-build:
	go build -o ./bin/app cmd/app/main.go

goose-build:
	go build -o ./bin/goose cmd/goose/main.go

# Used for hot reloading with air
air-build-dev: bun-build templ-build air-server-build goose-build

build-dev: bun-build templ-build dev-server-build goose-build

test:
	go test -v -cover ./...

test-coverage:
	[ -d coverage/ ] || mkdir -p coverage/
	go test -coverprofile=coverage/cover.out ./...
	go tool cover -html=coverage/cover.out -o=coverage/index.html
