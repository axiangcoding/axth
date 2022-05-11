.PHONY: clean upgrade check test

clean:
	go clean

upgrade:
	go get -u ./

check:
	go fmt ./
	go vet ./

test:
	go test ./...