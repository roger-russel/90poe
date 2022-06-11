.PHONY: run
run:
	go run ./cmd/90poe/main.go -file ./_test/assets/ports.json -buffer-size 4096

.PHONY: test
test:
	@go test --race ./... 

.PHONY: lint
lint:
	@golangci-lint run