set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

imageName := "leviathan:dev"

# default command
list:
    just --list

# build kraken
[working-directory: 'kraken']
krd:
	docker build . -t kraken:dev

krn:
    just krd
    docker run --rm --network=host kraken:dev

dk:
	docker build . -t {{imageName}}

lrn:
    just dk
    docker run --rm --network=host -v /var/run/docker.sock:/var/run/docker.sock {{imageName}}

# docker compose up
lrc:
    docker compose --profile lev up --build


# build leviathan with version and other metadata
dkv:
    docker build \
        --build-arg VERSION=test \
        --build-arg COMMIT_INFO=test \
        --build-arg BUILD_DATE=test \
        --build-arg BRANCH=test \
        -t {{imageName}} .


# docker compose up shorthand
up:
    docker compose up --build

# docker compose down shorthand
down:
    docker compose down

# start required services for levithan
dev:
    docker compose --profile dev up

alias dc := dclean
dclean:
    docker rm -f $(docker ps -aq)
    docker image prune -ay

dkrn:
	docker compose up --build

post:
    docker compose --profile post up

# update all go deps
[working-directory: 'src']
get:
    go get -v -u all

# lint go files
[working-directory: 'src']
lint:
	golangci-lint run

# go mod tidy
[working-directory: 'src']
tidy:
    go mod tidy

[working-directory: 'src']
vet:
    go vet ./...

# go build and run
[working-directory: 'src']
gb:
    go build -o ../bin/leviathan.exe

# go build
gr:
    just gb
    ./bin/leviathan.exe
