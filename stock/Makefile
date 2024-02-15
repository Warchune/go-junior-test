CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=go-junior-test/stock/cmd/app

all: goimports fmt build test

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

precommit: goimports fmt build test
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

fmt:
	go fmt ./...

goimports: install-imports
	goimports -w .

install-imports:
	go install golang.org/x/tools/cmd/goimports@latest

sql-build:
	../bin/sqlc generate