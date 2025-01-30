set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

dk:
	docker build . -t leviathan:dev

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

bdrn:
	docker compose up --build