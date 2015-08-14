export GOPATH=$(shell pwd)/../

build:
	go build -o ./webadmin ./server
get:
	cd ./server && go get -v -d
install:
	cd ./server && go install ./...