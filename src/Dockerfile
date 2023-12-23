## ビルドステージ
FROM --platform=linux/amd64 golang:1.21-alpine AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build \
    -o /go/bin/api_service \
    -ldflags '-s -w'
CMD ["/app/api_service"]
