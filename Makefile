init-pre-commit:
# Check https://pre-commit.com/ for more installation methods
	brew install pre-commit
	pre-commit install
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
