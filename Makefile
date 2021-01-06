BIN_PATH=bin
BIN_SOURCE=$(BIN_PATH)/web-crawler


.PHONE: build
build:
	# https://stackoverflow.com/a/21164441
	mkdir -p $(BIN_PATH)
	make test
	go build -modfile go.mod -o $(BIN_SOURCE) -a cmd/main.go

.PHONE: dev-run
dev-run:
	 go run cmd/main.go

.PHONE: test
test:
	go test ./internal/pkg -v
