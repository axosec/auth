VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo "dev")
GIT_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

build:
	go build -ldflags "\
		-X 'main.Version=$(VERSION)' \
		-X 'main.GitCommit=$(GIT_COMMIT)' \
		-X 'main.BuildDate=$(BUILD_DATE)'" \
		-o axosec-auth ./cmd/api

start:
	./axosec-auth

run: build start

docker-build:
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		-t axosec-auth:local .

