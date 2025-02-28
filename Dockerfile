FROM golang:1.24-alpine AS builder

# arg substitution
# https://stackoverflow.com/questions/44438637/arg-substitution-in-run-command-not-working-for-dockerfile
ARG VERSION
ENV BV=${VERSION}

# for sqlite
RUN apk update && apk add --no-cache gcc musl-dev
ENV CGO_ENABLED=1

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY ./src .

# build optimized binary without debugging symbols
RUN go build -ldflags "-s -w" -o app

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/app .

ENV IS_DOCKER=true

EXPOSE 9221

CMD ["./app"]
