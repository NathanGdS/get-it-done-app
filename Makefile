run/cli:
	@go run ./cmd/cli.go $(ARG)

build/cli:
	@go build -o ./bin/cli.exe ./cmd/cli.go