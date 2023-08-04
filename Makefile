COVERAGE_FILE = coverage.txt

.PHONY: test
test123:  ## Run go unit tests
	mkdir -p bin
	go test -coverprofile $(COVERAGE_FILE) ./...
	@echo "-- Test coverage --"
	@go tool cover -func=$(COVERAGE_FILE)|grep "total:"

test:
	mkdir -p bin
	go test -short -coverprofile=bin/cov.out `go list ./... | grep -v vendor/`
	go tool cover -func=bin/cov.out