.PHONY:all
all: build

build:
	go build -o bin/knowledge-map main.go

test:
	go test ./... -v -count=1 -coverprofile=coverage.txt -covermode=atomic

deploy:
	@ssh ${MY_KNOWLEDGE_MAP_HOST} \
		"cd ~/knowledge-map; \
		 git pull; \
		 echo ${MY_KNOWLEDGE_MAP_HOST_PASSWORD} | sudo -S systemctl restart knowledge-map"
