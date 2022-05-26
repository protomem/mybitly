
.PHONY: build
build:
	go build -o ./build/ ./cmd/api

.PHONY: run
run:
	./build/api

