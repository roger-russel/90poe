.PHONY: run
run:
	./scripts/env.sh go run ./cmd/90poe/main.go -f _test/assets/ports.json

.PHONY: test
test:
	@go test --race ./... 

.PHONY: lint
lint:
	@golangci-lint run