VERSION=0.1.0
LDFLAGS=-ldflags "-X main.version=`date -u +${VERSION}.build.%Y%m%d.%H%M%S`"
BUILDFLAGS=
NAME=wallet2020
MAIN=cmd/${NAME}/main.go

all: app

app:
	go build -race ${BUILDFLAGS} ${LDFLAGS} -o ${NAME} ${MAIN}

dist:
	go build ${BUILDFLAGS} ${LDFLAGS} -o ${NAME} ${MAIN}

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${NAME} ${BUILDFLAGS} ${LDFLAGS} ${MAIN}

clean:
	go clean -x ${MAIN}
	rm -f ${NAME}

test: stopdb startdb
	./scripts/test


lint:
	golint internal/... cmd/...

vet:
	go vet ${BUILDFLAGS} ./internal/app/...
	go vet ${BUILDFLAGS} ./cmd/...

startdb:
	docker-compose -p ${NAME}-dbonly -f test/docker/docker-compose.yml --log-level ERROR up -d

stopdb:
	docker-compose -p ${NAME}-dbonly -f test/docker/docker-compose.yml down

build-docker:
	scripts/build-docker.sh ${NAME}

.PHONY: app linux clean test lint docker publish
