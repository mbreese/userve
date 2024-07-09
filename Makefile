all: bin/userve bin/userve.darwin_amd64

bin/userve: userve/main.go
	go build -o bin/userve userve/main.go

bin/userve.darwin_amd64: bin/userve
	GOOS=darwin GOARCH=amd64 go build -o bin/userve.darwin_amd64 userve/main.go

.PHONY: all
