include .env
export

build-it:
	go build -o ./build/dynu-updater ./cmd
run:
	go run ./cmd/updater.go  
d-build:
	docker build -t iben12/dynu-updater .
d-run:
	docker run --rm --env-file .env iben12/dynu-updater