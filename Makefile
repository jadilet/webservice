
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/gorilla/mux
	
.PHONY: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o bin/server cmd/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t server:0.0.1
