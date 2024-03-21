.PHONY: build run clean

build: go build -o go-exercise ./cmd/go-exercise/app.go

run: go run ./cmd/go-exercise/app.go

clean: rm -f
