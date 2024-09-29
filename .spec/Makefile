gensrv:
	openapi-generator-cli generate -i ./leviathan.yaml -g go-gin-server -o ./server -c ./config-go-server.yml --git-repo-id makeopensource --git-user-id leviathan
	make mvapi

genclient:
	openapi-generator-cli generate -i ./leviathan.yaml -g typescript-axios -o ./client -c ./config-ts-client.yml --git-repo-id makeopensource --git-user-id leviathan

mvapi:
	cp -rf server/generated-server ../internal/
