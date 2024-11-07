# Go言語の静的解析のお勉強用

## なにから勉強したら良いか？

Go言語の静的解析ツールを作る場合、普通はgo/analysisパッケージを用いて作る。

### 導入的なやつ

- [(読み物)「静的解析とコード生成」スライド](https://docs.google.com/presentation/d/1I4pHnzV2dFOMbRcpA-XD0TaLcX6PBKpls6WxGHoMjOg/edit?usp=sharing)
  - => これが一番オススメ
- [(読み物)逆引き Goによる静的解析](https://zenn.dev/tenntenn/books/d168faebb1a739/viewer/22e4d4)
  - => これはオススメだけど、一部、ページが欠落してたりする。
- [go/analysisパッケージのgodoc](https://pkg.go.dev/golang.org/x/tools@v0.26.0/go/analysis)
- [(動画)2024/11/02 静的解析Night](https://www.youtube.com/watch?v=oBgDdx8gNQY)
- [(ツール)静的解析ツールのスケルトンコードを自動生成するツール](https://github.com/gostaticanalysis/skeleton)

## 構文解析

構文解析とはAST(抽象構文木)を解析すること。
ASTを走査し何らかの解析結果を返すこと。

Go言語において、[go/ast](https://pkg.go.dev/go/ast)パッケージがASTの生成器を提供する。
プログラマは`Parse`なんちゃら関数を呼ぶだけでASTを生成できる。

構文解析は、Linterを書くためのほぼほぼ必須知識。有用なLinterを作りたいのであれば、ほぼほぼ構文解析が必要となると思う。

初学者にとっては、構文解析の理解が最初のハードルとなると思われる。
ここについての良い資料ないだろうか？

[Go言語仕様](https://go.dev/ref/spec)を頭に入れておくと、空気を吸うようにASTが扱えるようになるのかもしれない。
でもこれはちょっとハードルが高い。
言語仕様の読解は、言語仕様の量が多いので時間がかかる。まぁでも量が多いだけで難しくはないので、時間があるなら言語仕様読むのはありかもしれん。

鉄板は、tenntennさんの「静的解析とコード生成」スライドの「構文解析」。ぶっちゃけコレ読むだけでいいのかもしれない。

その他、構文解析を助けてくれそうなツール

- [AST Viewer](https://yuroyoro.github.io/goast-viewer/)

## 型チェック

よく知らん。ということで勉強中。。。

[go/types](https://pkg.go.dev/go/types)

鉄板は、tenntennさんの「静的解析とコード生成」スライドの「型チェック」。ぶっちゃけコレ読むだけでいいのかもしれない。

## すでにあるソースコード読んでみる

自分で作るよりも、すでにあるものを読むほうが早い。
go/analysisパッケージを用いて書かれたLinterの例

- [golang.org/x/tools/go/analysis/passes 配下にあるLinter](https://cs.opensource.google/go/x/tools/+/refs/tags/v0.26.0:go/analysis/passes/)
- [golangci-lintで提供されているLinterたち](https://github.com/golangci/golangci-lint/tree/master/pkg/golinters)

# Memo

## linterのスケルトンコード生成

```bash
NAME=linter01 make linter
```

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

