
BINARY=canvas-cc

test: 
	go test -v ./...

app:
	go build -o dist/${BINARY} main.go
	cp settings.yml dist/settings.yml

unittest:
	go test -short  ./...

.PHONY: all test clean