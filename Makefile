dkrn:
	docker run --rm -p 9221:9221 leviathan:dev

bdrn:
	make dk
	make dkrn
buildrun:
	make dkbuild
	make dkrun

pullrun:
	docker run --rm --name leviathan -p 9221:9221 ghcr.io/makeopensource/leviathan:beta
