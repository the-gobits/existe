.PHONY: test open-cover

coverage.out: existe.go existe_test.go
	@rm -f coverage.out
	go test -v -coverprofile=coverage.out ./...

open-cover: coverage.out
	go tool cover -html=coverage.out

test: coverage.out
	@echo latest tests ran
