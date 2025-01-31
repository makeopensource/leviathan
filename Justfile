set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

imageName := "leviathan:dev"

dk:
	docker build . -t {{imageName}}

bdrn:
    just dk
    docker run --rm -v /var/run/docker.sock:/var/run/docker.sock {{imageName}}

dkrn:
	docker compose up

# lint go files
[working-directory: 'src']
lint:
	golangci-lint run

# go mod tidy
[working-directory: 'src']
tidy:
    go mod tidy
