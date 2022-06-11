# 90poe test

## Development

### Requirements

* [Golang](https://go.dev) 1.17 or higher;
* [Golanglint-ci](https://golangci-lint.run/) 1.46.2 or higher;
* [Make](https://man7.org/linux/man-pages/man1/make.1.html), actually is is not required for development, but the commands used on this project for development and run in production will be writen and versioned with on it Makefile.
* [Docker](https://www.docker.com/) 20.10.16 or higher.
* [Docker Compose](https://docs.docker.com/compose/) 20.10.16 or higher.

### Commands

Running local

```sh
make run # running // go run ./cmd/90poe/main.go -file ./_test/assets/ports.json -buffer-size 4096 
```

```sh
make cover # running test
```

```sh
make lint # running lint
```

Running Inside Docker

```sh
make docker/build
```

```sh
make docker/run
```

The docker files are inside [./build](./build) folder.
To change the json file that is running inside docker is necessary to:

1. Change the ENV FILE to the new file
2. Change the volume mounted or put it inside _test/assets folder

## Project Decisions and Considerations

* At first I tried to find a fine library to parse huge JSONs but this type of JSON is quite different from the established patterns like ND JSON, because I didn't find any I had to write my on Parser, it was fun, so I finished it even if the time passed a little.
* I implemented a DB component to look more like how I would do dependency injection.
* There are a lot of things that I would need to do to make my Streamer components handle any JSON, but I just focus on making it handle the desired one.
* Makefile is life, if you don't have the Makefile installed on your machine please, take a look inside Makefile to see the commands there.
* Normally for development I would use VSCode Remote Containers Plugin, but because I don't know who will take a look in this project is using the VSCode, I will not use it on this project.
* I did a few tests to show how I do tests, and not to get high coverage, but because I did a BDD test the coverage is much higher than shown on the `go test` command, to see the real coverage use `make cover` and see into coverage.html file by file, it will be higher there because the command into Makefile will get the coverage over all packages instead of just on the package that the test is running.
* I did Dockerfile and Docker-Compose only for Build and Production, for development I would use Remote Containers as I said before, but I didn't set up it on this project.
* For Lint I prefer to use Golanglint-ci because it has a lot of good Lints on it.
* For tests I just did test table-driven test, it is just my personal preference, but I don't have a problem in adapting to use the company standards.
