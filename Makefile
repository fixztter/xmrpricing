build:
	@go build -o bin/monero-price cmd/main.go

run: build
	@./bin/monero-price

test:
	@go test -v ./...

clean:
	@rm -rfv ./bin
