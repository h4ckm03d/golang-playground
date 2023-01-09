build:
	rm -rf api
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o api . 
build-docker:
	docker build -t hello-go .
