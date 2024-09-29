FROM golang:1.19 AS build
WORKDIR /go/src
COPY generated-server ./generated-server
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o serverstub .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/serverstub ./
EXPOSE 8080/tcp
ENTRYPOINT ["./serverstub"]
