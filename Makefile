
BINARY=canvas-cc

test: 
	go test -v ./...

app:
	go build -o dist/${BINARY} main.go
	cp settings.yml dist/settings.yml

test-short:
	go test -short  ./...

.PHONY: all test clean