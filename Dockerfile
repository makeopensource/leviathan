FROM golang:1.24-alpine AS builder

# for sqlite
RUN apk update && apk add --no-cache gcc musl-dev
ENV CGO_ENABLED=1

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY ./src .

# Build arguments
ARG VERSION=dev
ARG COMMIT_INFO=unknown
ARG BUILD_DATE=unknown
ARG BRANCH=unknown

# arg substitution, do not put it higher than this for caching
# https://stackoverflow.com/questions/44438637/arg-substitution-in-run-command-not-working-for-dockerfile
ENV VERSION=${VERSION}
ENV COMMIT_INFO=${COMMIT_INFO}
ENV BUILD_DATE=${BUILD_DATE}
ENV BRANCH=${BRANCH}

# build optimized binary without debugging symbols
RUN SOURCE_HASH=$(find . -type f -name "*.go" -print0 | sort -z | xargs -0 cat | sha256sum | cut -d ' ' -f1) && \
    go build -ldflags "-s -w \
      -X github.com/makeopensource/leviathan/internal/info.Version=${VERSION} \
      -X github.com/makeopensource/leviathan/internal/info.CommitInfo=${COMMIT_INFO} \
      -X github.com/makeopensource/leviathan/internal/info.BuildDate=${BUILD_DATE} \
      -X github.com/makeopensource/leviathan/internal/info.Branch=${BRANCH} \
      -X github.com/makeopensource/leviathan/internal/info.SourceHash=${SOURCE_HASH}" \
    -o leviathan  \
    ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/leviathan .

ENV LEVIATHAN_IS_DOCKER=true
# default level info when running in docker
ENV LEVIATHAN_LOG_LEVEL=info

EXPOSE 9221

CMD ["./leviathan"]
