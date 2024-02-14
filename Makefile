build-all:
	cd stock && GOOS=linux GOARCH=amd64 make build

up: build-all
	docker-compose up --force-recreate --build

precommit:
	cd stock && make precommit
