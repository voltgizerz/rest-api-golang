run:
	go run main.go

format:
	gofmt -w ./..
	revive ./...

build:
	go build