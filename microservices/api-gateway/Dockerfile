FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN apk update && \
    apk add build-base

RUN go mod download

COPY . .

RUN go build -o medview main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/medview /app/medview

CMD ["./medview", "api-gateway"]

EXPOSE 9000


