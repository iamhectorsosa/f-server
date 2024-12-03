default: dev

format:
	@go fmt ./...

test:
	@go test ./...

dev/server:
	go run github.com/air-verse/air@latest \
	--build.cmd "go build -o tmp/main" \
	--build.bin "./tmp/main" \
	--build.delay "0" \
	--build.include_ext "go" \
	--misc.clean_on_exit true

dev:
	@make dev/server
