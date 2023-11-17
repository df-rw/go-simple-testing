.PHONY: run test coverage lint vet

run: lint vet
	go run ./main.go

test: lint vet
	go test -v ./...

coverage: test
	go tool cover -html=./coverage.out

lint:
	golint ./...

vet:
	go vet ./...

clean:
	rm -f coverage.out
