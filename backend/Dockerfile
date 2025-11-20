FROM golang:1.23.1-alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /app
# Cache deps
COPY go.mod go.sum ./
RUN go mod download
# Copy the source
COPY . .
# Build server binary (cgo para sqlite)
ENV CGO_ENABLED=1 GOOS=linux
RUN go build -o server ./cmd/server

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata libstdc++ libc6-compat
WORKDIR /root/
COPY --from=builder /app/server /root/server
EXPOSE 8080
ENTRYPOINT ["/root/server"]