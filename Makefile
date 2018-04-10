COMMIT = $$(git describe --always)

generate:
	@go generate ./...

build: generate
	@echo "====> Build tools"
	@go build -o ./bin/tools main.go

build.cli: generate
	@echo "====> Build tools cli"
	@go build -ldflags "-X github.com/zcong1993/tools/cmd/main.GitCommit=\"$(COMMIT)\"" -o ./bin/tools ./cmd/main.go
