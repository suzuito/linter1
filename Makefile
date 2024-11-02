.PHONY: linter1
linter1-run_on_govet:
linter1.vettool.out:
	go build -o $@ linter1/cmd/vettool/*.go