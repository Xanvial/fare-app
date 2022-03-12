build:
	@go build -v -o fare-app cmd/*.go

run: build
	@./fare-app

gen-mock:
	@go generate ./...

test:
	@go test ./...

test-cover:
	@go test ./... -coverprofile cover.out
	@go tool cover -func=cover.out

build-all:
	@env GOOS=darwin GOARCH=amd64 go build -v -o bin/mac/fare-app cmd/*.go
	@env GOOS=linux GOARCH=amd64 go build -v -o bin/linux/fare-app cmd/*.go
	@env GOOS=linux GOARCH=386 go build -v -o bin/linux_x86/fare-app cmd/*.go
	@env GOOS=windows GOARCH=amd64 go build -v -o bin/win/fare-app.exe cmd/*.go
