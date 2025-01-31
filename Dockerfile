FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ./src .

# get depedencies
RUN go mod tidy

# build optimized binary without debugging symbols
RUN go build -ldflags "-s -w" -o app cmd/leviathan-agent/main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/app .

ENV IS_DOCKER=true

EXPOSE 9221

CMD ["./app"]
