run:
	@ PROXY_GEO_CONFIG=internal/repository/config/config.yml go run cmd/cli/main.go

lint:
	@ golangci-lint run

it:
	@ docker compose up -d
	@ go test -count=1 -v ./internal/test
	@ docker compose down

build:
	@ GOOS=linux GOARCH=amd64 go build -o checkproxy cmd/cli/main.go