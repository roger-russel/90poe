# 90poe test

## Development

### Requirements

* [Golang](https://go.dev) 1.17 or higher;
* [Golanglint-ci](https://golangci-lint.run/) 1.46.2 or higher;
* [Make](https://man7.org/linux/man-pages/man1/make.1.html), actually is is not required for development, but the commands used on this project for development and run in production will be writen and versioned with on it Makefile.
* [Docker](https://www.docker.com/) 20.10.16 or higher.
* [Docker Compose](https://docs.docker.com/compose/) 20.10.16 or higher.

### Commands

```sh
make run # go run cmd/90poe/main.go
```

```sh
make test # go test -v --race ./...
```

```sh
make lint # golangci-lint run
```

## Project Decisions

* In Memory Database will be used [DragonflyDB](https://dragonflydb.io/), it is a lightweight and fast in-memory database.
