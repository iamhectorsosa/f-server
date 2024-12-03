# Game Server Test Repository

## Introduction

This is a test repository to model an HTTP Server (using `stdlib`) in Go using the project structure:

```bash
├── Makefile                 # commands and live-reload server
├── config                   # env vars
├── internal
│   ├── database             # database connections and inits
│   │   ├── queries          # sqlc-generated files
│   │   └── sql
│   │       ├── migrations   # sql migration files (incl. schema)
│   │       └── queries      # sql queries for sqlc generation
│   ├── server               # server, handlers, and middleware
│   └── store                # store for db interactions and business logic
├── main.go
└── sqlc.yaml
```

## Requirements

- Go version `1.23.2`
- Goose >= v3 (for database migrations)
- sqlc >=v1.27.0 (for sql generation)

## TODO

- [ ] Improve Store, server and auth error handling
- [ ] Integrate a logger into the Server
- [ ] Integrate business logic to improve metrics on matches
- [ ] Optimize SQL schema and/or queries with redesign or indexesa
- [ ] Add an integration test to the Server
- [ ] Add github workflow to check tests, staticcheck, gosec
- [ ] Add Docker
