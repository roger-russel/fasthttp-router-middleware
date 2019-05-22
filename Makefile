.PHONY: test coverage

test:
	@go test ./... -cover -coverprofile=./system.cov

coverage: test
	@go tool cover -html=./system.cov -o coverage.html
	@google-chrome coverage.html
