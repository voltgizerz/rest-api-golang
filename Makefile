run:
	go run main.go

format:
	gofmt -w ./..
	revive ./...

build:
	go build

mockgen:
	mockgen -source=./service/pokemon.go -destination ./mocks/mocks_servicePokemon.go
	mockgen -source=./repository/pokemon.go -destination ./mocks/mocks_repoPokemon.go 

test:
	go test -v -cover ./...

cover-out:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out