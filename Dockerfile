FROM golang:alpine AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["air", "-c", ".air.toml"]

FROM dev AS builder
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o mindbox-api ./cmd/server/server.go

FROM alpine:latest as prod
WORKDIR /root/
COPY --from=builder /app/mindbox-api .
CMD ["./mindbox-api"]