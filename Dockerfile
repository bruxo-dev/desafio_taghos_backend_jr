FROM golang:alpine AS builder

RUN apk update && apk add --no-cache ca-certificates curl bash

WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,https://goproxy.io,https://gocenter.io,direct

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/db/init.sql ./db/

EXPOSE 8080

CMD ["./main"]
