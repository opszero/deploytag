.PHONY=build
push:
	docker build -t opszero/deploytag:v3 .
	docker push opszero/deploytag:v3

