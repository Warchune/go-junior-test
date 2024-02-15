build-all:
	cd stock && GOOS=linux GOARCH=amd64 make build
	cp bin/goose stock/bin/goose

up: build-all
	docker-compose up --force-recreate --build

precommit:
	cd stock && make precommit

install-deps:
	GOBIN=$(CURDIR)/bin go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1
	GOBIN=$(CURDIR)/bin go install github.com/pressly/goose/v3/cmd/goose@latest