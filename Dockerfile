FROM golang:alpine3.20 AS builder

WORKDIR /chat-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /chat-service/main .

ENV PORT=8082

ENTRYPOINT [ "./main" ]
