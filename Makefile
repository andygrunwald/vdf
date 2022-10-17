.DEFAULT_GOAL := help

.PHONY: help
help: ## Outputs the help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: staticcheck
staticcheck: ## Runs static code analyzer staticcheck
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

.PHONY: vet
vet: ## Runs go vet
	go vet ./...

.PHONY: test
test: ## Runs all unit tests
	go test -v -race ./...

.PHONY: test-coverage
test-coverage: ## Runs all unit tests + gathers code coverage
	go test -v -race -coverprofile coverage.txt ./...

.PHONY: test-coverage-html
test-coverage-html: test-coverage ## Runs all unit tests + gathers code coverage + displays them in your default browser
	go tool cover -html=coverage.txt

.PHONY: test-fuzzing
test-fuzzing: ## Runs all unit fuzzing tests (each test with a timeout)
	go test -fuzz=FuzzScanner_ScanWithoutWhitespace -fuzztime 45s
	go test -fuzz=FuzzScanner_ScanWithWhitespace -fuzztime 45s
	go test -fuzz=FuzzParser_Parse -fuzztime 45s