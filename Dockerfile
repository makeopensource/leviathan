FROM golang:1.24-alpine AS builder

# Build arguments
ARG VERSION=dev
ARG COMMIT_INFO=unknown
ARG BUILD_DATE=unknown
ARG BRANCH=unknown

# for sqlite
RUN apk update && apk add --no-cache gcc musl-dev
ENV CGO_ENABLED=1

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY ./src .

# arg substitution, do not put it higher than this for caching
# https://stackoverflow.com/questions/44438637/arg-substitution-in-run-command-not-working-for-dockerfile
ENV VERSION=${VERSION}
ENV COMMIT_INFO=${COMMIT_INFO}
ENV BUILD_DATE=${BUILD_DATE}
ENV BRANCH=${BRANCH}

# build optimized binary without debugging symbols
RUN go build -ldflags "-s -w \
      -X github.com/makeopensource/leviathan/common.Version=${VERSION} \
      -X github.com/makeopensource/leviathan/common.CommitInfo=${COMMIT_INFO} \
      -X github.com/makeopensource/leviathan/common.BuildDate=${BUILD_DATE} \
      -X github.com/makeopensource/leviathan/common.Branch=${BRANCH}" \
    -o leviathan

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/leviathan .

ENV IS_DOCKER=true

EXPOSE 9221

CMD ["./leviathan"]
