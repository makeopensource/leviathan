dkbuild:
	docker build . -t github.com/makeopensource/leviathan

dkrun:
	docker run -p 9221:9221 github.com/makeopensource/leviathan

buildrun:
	make dkbuild
	make dkrun

pullrun:
	docker run -p 9221:9221 ghcr.io/makeopensource/leviathan