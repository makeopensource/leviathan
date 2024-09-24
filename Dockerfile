FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update

COPY . .

# build optimized binary without debugging elements
RUN go build -ldflags "-s -w" -o app cmd/leviathan-agent/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

EXPOSE 9221

CMD ["./app"]
