build:
	@go build -o bin/uuid main.go
build-arm:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o bin/uuid main.go
run: build
	@./bin/uuid