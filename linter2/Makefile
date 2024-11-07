.PHONY: vettool
vettool: vettool.out
	go vet -vettool=$(abspath $<) ./testdata/src/a/a.go

vettool.out: $(shell find . -name "*.go")
	go build -o $@ cmd/vettool/*.go

test:
	go test ./...
