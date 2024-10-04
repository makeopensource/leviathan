dk:
	docker build . -t github.com/makeopensource/leviathan

dkrun:
	docker run -p 9221:9221 github.com/makeopensource/leviathan

buildrun:
	make dk
	make dkrun