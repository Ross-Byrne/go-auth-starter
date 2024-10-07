# go-auth-starter
A starter go project with auth. Uses Echo web server, Templ and HTMX for frontend, Sqlite3 and Goose for DB migrations and Air for hot reloading.

## Setup

### Mise
Install [Mise](https://mise.jdx.dev/) to manage dependances for go and bun.

Run `mise install` to install or `mise up` to update

### Install go dependances
There are a number of go tools required, such as LSP for Go and Templ. Run the following:
```bash
go install golang.org/x/tools/gopls@latest && \
go install github.com/a-h/templ/cmd/templ@latest && \
go install github.com/air-verse/air@latest && \
go install github.com/pressly/goose/v3/cmd/goose@latest
```


### ENV Variables
create `.env` file with the following:

```
PRODUCTION=false
GOOSE_DRIVER=sqlite3
GOOSE_DBSTRING=<db_string>
GOOSE_MIGRATION_DIR=./database/migrations
ASSETS_PATH=./app/assets
SESSION_SECRET_KEY=<generate your own secret key. 32 or 64 chars long>
```

### Build Project
Use Make to build the project. Run `make build-dev`.

During development, use Air for hot reloading. Run `air` in project directory should build and start the webserver.
