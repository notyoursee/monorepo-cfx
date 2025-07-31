FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o goapp

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/goapp .

EXPOSE 8080
CMD ["./goapp"]
