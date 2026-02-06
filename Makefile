.PHONY: test
test:
	go clean -testcache
	CGO_ENABLED=0 go test -v ./test/...
