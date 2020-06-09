BUILD_ENV := CGO_ENABLED=0
BUILD=`date +%FT%T%z`
VERSION=0.0.1
# latest tag
TAG=$(git describe --tags `git rev-list --tags --max-count=1`)
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Tag=${TAG} -X main.Build=${BUILD}"

TARGET_EXEC := mars_agent

.PHONY: all clean setup mod build-linux build-osx

all: clean setup mod build-linux build-osx


mod:
	go mod tidy
	go mod vendor

clean:
	rm -rf ./bin

setup:
	mkdir -p ./bin/linux
# 	mkdir -p ./bin/osx

build-linux: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -mod=vendor -o ./bin/linux/${TARGET_EXEC}

# build-osx: setup
# 	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -mod=vendor -o ./bin/osx/${TARGET_EXEC}