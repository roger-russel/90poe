# syntax=docker/dockerfile:1

# Prod
## Builder binary for Prod
FROM golang:1.17-alpine AS builder

WORKDIR /build
RUN apk add --no-cache git make

COPY go.mod go.sum /
RUN go mod download

COPY . .
RUN make build
RUN chmod +x /build/bin/app

## Runner image for prod
FROM alpine as runner
USER nobody

ENV FILE=""
ENV BUFFER_SIZE=""

WORKDIR /

COPY --from=builder /build/bin/app .
ENTRYPOINT ./app -file ${FILE} -buffer-size ${BUFFER_SIZE}
