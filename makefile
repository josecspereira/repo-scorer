.PHONY: run
run:
	go run src/main.go

.PHONY: test
test:
	go test -v -race -buildvcs ./...