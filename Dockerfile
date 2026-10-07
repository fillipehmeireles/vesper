FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /vesper ./cmd/main.go


FROM alpine:latest

COPY --from=builder /vesper /usr/local/bin/vesper

ENTRYPOINT ["vesper"]
