.PHONY: clean open-cover test test-with-coverage

clean:
	@rm -f coverage.out

coverage.out: existe.go existe_test.go
	@rm -f coverage.out
	@go test -v -coverprofile=coverage.out ./...

open-cover: coverage.out
	@go tool cover -html=coverage.out

test:
	@go test -v -cover ./...

test-with-coverage: coverage.out
	@echo ran coverage generating test
