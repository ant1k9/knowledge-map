.PHONY:all
all: build

build:
	go build -o bin/app main.go
