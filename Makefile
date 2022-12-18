.PHONY: build
build:
	docker-compose build sima-app

run:
	docker-compose up sima-app

test:
	go test -v -timeout 30s ./...

run: build
