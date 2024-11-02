# Go言語の静的解析

Go言語の静的解析ツールを作る場合、普通はgo/analysisパッケージを用いて作る。

## なにから勉強したら良いか？

- [(読み物)静的解析とコード生成](https://docs.google.com/presentation/d/1I4pHnzV2dFOMbRcpA-XD0TaLcX6PBKpls6WxGHoMjOg/edit?usp=sharing)
  - => ざっくりと概略掴みたいだけならこっちがおすすめ
- [(読み物)逆引き Goによる静的解析](https://zenn.dev/tenntenn/books/d168faebb1a739/viewer/22e4d4)
  - => 詳細に書いてあるのが好きならこっちがおすすめ
- [go/analysisパッケージのgodoc](https://pkg.go.dev/golang.org/x/tools@v0.26.0/go/analysis)
- [(動画)2024/11/02 静的解析Night](https://www.youtube.com/watch?v=oBgDdx8gNQY)
- [(ツール)静的解析ツールのスケルトンコードを自動生成するツール](https://github.com/gostaticanalysis/skeleton)

## go/analysisパッケージを用いて書かれたLinterの例

- [golang.org/x/tools/go/analysis/passes 配下にあるLinter](https://cs.opensource.google/go/x/tools/+/refs/tags/v0.26.0:go/analysis/passes/)
- [golangci-lintで提供されているLinterたち](https://github.com/golangci/golangci-lint/tree/master/pkg/golinters)

# Memo

### ソースコード



### 実行方法

オプションを何も指定せずに`go vet`コマンドを実行したとき、passes配下にあるLinterが適用される。

```bash
% go vet ./testdata/src/a/...        
# github.com/suzuito/linter1/testdata/src/a
testdata/src/a/a.go:6:6: append with no values
testdata/src/a/a.go:11:2: unreachable code
```

`-appends`など指定すると、golang.org/x/tools/go/analysis/passes/appends のみが適用される。

```bash
% go vet -appends ./testdata/src/a/...
# github.com/suzuito/linter1/testdata/src/a
testdata/src/a/a.go:6:6: append with no values
```

## [golang.org/x/tools/go/analysis/passes](https://github.com/golangci/golangci-lint/tree/master/pkg/golinters) 配下にあるLinter

