FROM golang:1.24.5-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o testTaskEM .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/testTaskEM .

EXPOSE 8000

CMD ["./testTaskEM"]