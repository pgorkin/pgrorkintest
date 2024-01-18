all: prepare lint

# Installs all dependencies
prepare:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

# Lints codebase
lint:
	${GOPATH}/bin/golangci-lint run