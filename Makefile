
BINARY=canvas-env-checker

test: 
	go test -v ./...

engine:
	go build -o dist/${BINARY} main.go
	cp settings.yml dist/settings.yml

unittest:
	go test -short  ./...