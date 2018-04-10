COMMIT = $$(git describe --always)

generate:
	@go generate ./...

build: generate
	@echo "====> Build tools"
	@go build -o ./bin/tools main.go

build.cli: generate
	@echo "====> Build tools cli"
	@go build -o ./bin/tools ./cmd/main.go
