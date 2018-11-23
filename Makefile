build:
	go build -o bin/jenkinsctl main.go
run:
	go run main.go
test:
	go test -v ./...
cover:
	go test -cover -v ./...
