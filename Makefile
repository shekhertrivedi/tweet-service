
COVERAGE_FILE = coverage/coverage.txt

.PHONY: test
test:  ## Run go unit tests
	mkdir -p coverage
	go test -coverprofile $(COVERAGE_FILE) ./...
	@echo "-- Test coverage --"
	@go tool cover -func=$(COVERAGE_FILE)|grep "total:"
