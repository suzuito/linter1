# analysis.Analyzerを用いて作られたlinterはgo vetから呼び出せる
.PHONY: vettool
vettool: vettool.out
	go vet -vettool=$(abspath $<) ./testdata/src/a/a.go
vettool.out: $(shell find . -name "*.go")
	go build -o $@ cmd/vettool/*.go

# golangciから呼び出す場合にはいくつか選択肢がある。

# (1) golangci-lintの公式linterとして追加する
# -> https://golangci-lint.run/contributing/new-linters/#how-to-add-a-public-linter-to-golangci-lint

# (2) Go Plugin System
# 公式linterとして追加せずにプライベートで利用するだけであれば、
# Go Plugin Systemを用いるという方法がある
# -> https://golangci-lint.run/plugins/go-plugins
# で、下記はGo Plugin(.soファイル)をビルドするコマンド
# linter2.so: $(shell find . -name "*.go")
# 	go build -buildmode=plugin -o linter2.so plugin/golangci-lint/main.go
# なのですが、Go Pluginを用いる方法はかなりダルく挫折

test:
	go test ./...
