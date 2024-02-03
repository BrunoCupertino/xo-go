runclient:
	@go run cmd/client/main.go 

buildclient:
	@go build -o bin/client cmd/client/main.go 

runserver:
	@go run cmd/server/main.go 

buildserver:
	@go build -o bin/server cmd/server/main.go

build: buildclient buildserver

tests:
	@go test ./... -coverprofile cover.out

testscoverage: tests
	@go tool cover -html cover.out -o cover.html	&& open cover.html