FROM golang:1.24.5-alpine3.22

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

ENTRYPOINT ["CompileDaemon", \
            "-build=go build -o /build/app", \
            "-command=/build/app", \
            "-directory=/app"]
