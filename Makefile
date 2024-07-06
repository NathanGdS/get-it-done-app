run/cli:
	@./bin/cli.exe --$(ARG)

build/cli:
	@go build -o ./bin/cli.exe ./cmd/cli.go