run: build start

test:
	@echo "Ambalwarsa >> running test ..."
	@go test -v -race ./...

build:
	@echo "Ambalwarsa >> building binaries ..."
	@go build -o bin/ambalwarsa app.go
	@echo "Ambalwarsa >> success"

start:
	@echo "Ambalwarsa >> starting binaries ..."
	@./bin/ambalwarsa