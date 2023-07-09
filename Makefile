run:
	go run ./cmd/main.go

format:
	gofmt -w ./..
	revive -config revive.toml ./...

build:
	go build ./cmd/main.go

mockgen:
	mockgen -source=./internal/app/service/pokemon.go -destination ./mocks/mocks_servicePokemon.go
	mockgen -source=./internal/app/repository/pokemon.go -destination ./mocks/mocks_repoPokemon.go 

test:
	go test -v -cover ./...

cover-out:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out