.PHONY: test coverage

# Run tests
test:
	go test ./internal/modules/... -v

# Show coverage UI in browser with highlighted uncovered lines
coverage:
	@echo "Generating coverage report..."
	@go test ./internal/modules/... -coverprofile=coverage.out
	@echo "Opening coverage report in browser..."
	@go tool cover -html=coverage.out