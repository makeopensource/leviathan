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

bdrn:
    just dk
    docker run --rm --network=host -v /var/run/docker.sock:/var/run/docker.sock -v appdata:/app/appdata/ {{imageName}}

dkrn:
	docker compose up --build

post:
    docker compose --profile post up

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
