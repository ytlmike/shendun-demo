BUILD_NAME:=demo
BUILD_VERSION:=1.0
SOURCE:=*.go
LDFLAGS:=-ldflags "-X main.Version=${BUILD_VERSION}"

all: deps build

deps:
	go mod tidy

build: deps
	go build -o ${BUILD_NAME} ${SOURCE}

clean:
	go clean
	rm demo test.db
