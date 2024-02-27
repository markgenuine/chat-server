FROM golang:1.22-alpine AS builder

COPY . /github.com/markgenuine/chat-server/source/
WORKDIR /github.com/markgenuine/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/server/main.go

FROM alpine:3.19.1

WORKDIR /root/
COPY --from=builder /github.com/markgenuine/chat-server/source/bin/chat_server .
COPY local.env .

CMD ["./chat_server", "-config-path", "local.env"]