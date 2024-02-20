runclient:
	@go run cmd/client/main.go

buildclient:
	@go build -o bin/xogo cmd/client/main.go

runserver:
	@go run cmd/server/main.go 

buildserver:
	@go build -o bin/xogo-server cmd/server/main.go

build: buildclient buildserver

test:
	@go test ./... -coverprofile cover.out

testscoverage: tests
	@go tool cover -html cover.out -o cover.html	&& open cover.html