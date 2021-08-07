.PHONY:all
all: build

build:
	go build -o bin/knowledge-map main.go
