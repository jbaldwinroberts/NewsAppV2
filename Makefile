PWD=$(shell pwd)

deps: tools mocks
	go mod download

tools:
	# Install mockgen
	go get github.com/golang/mock/mockgen

mocks: clean_mocks
	go generate -v ./...

clean_mocks:
	rm -rf mocks/

run:
	go run main.go

test:
	go test -timeout=5s -cover -race -v ./...