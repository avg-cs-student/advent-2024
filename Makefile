all: fmt tidy test

fmt:
	go fmt ./...

tidy:
	go mod tidy

test:
	go test -count=1 -v ./...
