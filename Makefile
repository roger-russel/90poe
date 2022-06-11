.PHONY: run
run:
	go run ./cmd/90poe/main.go -file ./_test/assets/ports.json -buffer-size 4096

.PHONY: test
test:
	@go test -coverpkg ./... --race -coverprofile coverage.out ./... 

.PHONY: cover
cover: test
	@go tool cover -html=coverage.out

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: build
build:
	@go build -o ./bin/app ./cmd/90poe

.PHONY: docker/build
docker/build:
	@docker build . -f build/Dockerfile --target builder
	@docker-compose -f build/docker-compose.yaml build

.PHONY: docker/run
docker/run:
	@docker-compose -f build/docker-compose.yaml up