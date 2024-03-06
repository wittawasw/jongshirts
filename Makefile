build:
	go build -o bin/jongshirtsweb cmd/jongshirts/main.go
run:
	go run cmd/jongshirts/main.go
test:
	go test ./...