.PHONY: test coverage

packages:
	@dep ensure

test:
	@go test ./... -cover -coverprofile=./coverage.txt -covermode=atomic

coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
	@google-chrome coverage.html
