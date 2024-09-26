FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update

COPY . .

# get depedencies
RUN go mod tidy

# build optimized binary without debugging symbols
RUN go build -ldflags "-s -w" -o app cmd/leviathan-agent/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 9221

CMD ["./app"]
