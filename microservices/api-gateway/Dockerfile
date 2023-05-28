FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN apk update && \
    apk add build-base

RUN go mod download

COPY . .

RUN go build -o medica main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/medica /app/medica

CMD ./medica api-gateway --core-address=$CORE --core-port=9010

EXPOSE 9000