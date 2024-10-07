dk:
	docker build . -t leviathan:dev

dkrn:
	docker compose up

lint:
	golangci-lint run

bdrn:
	docker compose up --build