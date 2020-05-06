.PHONY: test
.PHONY: lint

go-get:
	go get -u -t ./...

install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.25.0

lint:
	./bin/golangci-lint run

test:
	go test ./... -v -cover
	make -C section1 test
