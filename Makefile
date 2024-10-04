dk:
	docker build . -t leviathan:dev

dkrn:
	docker run --rm -p 9221:9221 leviathan:dev

bdrn:
	make dk
	make dkrn