include .env
export

build-it:
	go build -o ./build/dynu-updater ./cmd/...
run:
	go run ./cmd/...
d-build:
	docker build -t dynu-updater .
d-run:
	docker run --rm --env-file .env dynu-updater