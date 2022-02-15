FROM golang:1.17 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 G00S=linux go build -o app main.go

FROM alpine:latest AS dockerize
COPY --from=builder /app .
CMD ["./app"]