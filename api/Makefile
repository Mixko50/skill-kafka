image-name = registry.gitlab.com/apisit.m2/skill-api-kafka/skill-api

run:
	go run main.go

test:
	go test ./... -cover

build: test
	docker build -t $(image-name):latest -f ./Dockerfile .

push:
	docker login registry.gitlab.com -u $(shell bash -c 'read -p "Username: " username; echo $$username') -p $(shell bash -c 'read -s -p "Password: " pwd; echo $$pwd')
	docker push $(image-name):latest