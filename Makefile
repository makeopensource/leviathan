dk:
	docker build . -t leviathan:dev

dkrn:
	docker compose up --build

bdrn:
	make dk
	make dkrn