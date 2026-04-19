# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/api ./cmd/api

FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata curl && update-ca-certificates

COPY --from=builder /out/api /app/api

EXPOSE 8000

HEALTHCHECK --interval=15s --timeout=3s --start-period=20s --retries=5 \
  CMD curl -fsS http://127.0.0.1:${PORT:-8000}/healthz || exit 1

CMD ["/app/api"]
