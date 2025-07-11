FROM golang:1.24.1 AS builder

WORKDIR /usr/src/app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/bin/app /usr/local/bin/app
COPY --from=builder /go/bin/goose /usr/local/bin/goose

WORKDIR /app

COPY .env .
COPY wait-for-it.sh .
COPY ./sql/schema/ ./sql/schema
COPY ./ui ./ui

COPY entrypoint.sh entrypoint.sh
RUN chmod +x ./entrypoint.sh

CMD ["./entrypoint.sh"]
