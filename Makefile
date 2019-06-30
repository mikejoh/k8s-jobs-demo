# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

dockerize: build-sender build-logger

build-sender:
	docker build ./job-sender -t mikejoh/job-sender:latest; docker push mikejoh/job-sender

build-logger:
	docker build ./job-logger -t mikejoh/job-logger:latest; docker push mikejoh/job-logger
