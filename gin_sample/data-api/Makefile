BINARY_NAME=main

run: 
	make format
	make swag
	go run .

swag:
	swag init

format:
	swag fmt
	go fmt ./...

test:
	go test ./... -v

install:
	go get .

docker:
	docker build -t my-golang-app .

build:
	go mod tidy
	go build -o ${BINARY_NAME} main.go

deploy:
	docker buildx build --platform linux/amd64 -t data-api .
	docker tag data-api asia-east1-docker.pkg.dev/communityaimeeting/data-api/data-api
	docker push asia-east1-docker.pkg.dev/communityaimeeting/data-api/data-api

clean:
	go clean