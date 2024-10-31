build:
	@go build -o bin/xmrpricing ./cmd/xmrpricing/main.go

run: build
	@./bin/xmrpricing

test:
	@go test -v ./...

clean:
	@rm -rfv ./bin
