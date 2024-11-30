# Go言語の静的解析のお勉強用

## なにから勉強したら良いか？

Go言語の静的解析ツールを作る場合、普通はgo/analysisパッケージを用いて作る。

### 導入的なやつ

- [(読み物)「静的解析とコード生成」スライド](https://docs.google.com/presentation/d/1I4pHnzV2dFOMbRcpA-XD0TaLcX6PBKpls6WxGHoMjOg/edit?usp=sharing)
  - => これが一番オススメ
- [(読み物)逆引き Goによる静的解析](https://zenn.dev/tenntenn/books/d168faebb1a739/viewer/22e4d4)
  - => これもオススメだけど、一部、ページが欠落してたりする。
- [go/analysisパッケージのgodoc](https://pkg.go.dev/golang.org/x/tools@v0.26.0/go/analysis)
- [(動画)2024/11/02 静的解析Night](https://www.youtube.com/watch?v=oBgDdx8gNQY)
- [(ツール)静的解析ツールのスケルトンコードを自動生成するツール](https://github.com/gostaticanalysis/skeleton)

## 静的解析の流れ

- 字句解析する(go/ast がやる)
- 抽象構文木を作る(go/ast がやる)
- 抽象構文木をゴニョゴニョする(自作するところ)
- 型情報を割り当てる(go/types がやる)
- 型情報をゴニョゴニョする(自作するところ)

## 構文解析

構文解析とはAST(抽象構文木)を生成すること。

Go言語において、[go/ast](https://pkg.go.dev/go/ast)パッケージがASTの生成器を提供する。
静的解析ツール作成者は`Parse`なんちゃら関数を呼ぶだけでASTを生成できる。

静的解析ツール作成者は、「生成されたASTを探索し、ASTが意図通りでなければアサーションを上げる」というコードを書く。
これが静的解析ツール。

AST構造を理解しておくことが重要。
https://motemen.github.io/go-for-go-book/

[Go言語仕様](https://go.dev/ref/spec)を頭に入れておくと、空気を吸うようにASTが扱えるようになる。
が、これはちょっとハードルが高い気がする。まあ、Go言語仕様を読めるようになっておけば大丈夫。

その他、構文解析を助けてくれそうなツール

- [AST Viewer](https://yuroyoro.github.io/goast-viewer/)

## 型チェック

型チェックとは、ASTを入力として、ASTに対して型情報を割り当てること。

ASTがあれば、Goは型情報を割り当てることができる(らしい)。go/typesパッケージがそれをしている。
[go/types](https://pkg.go.dev/go/types)
(tenntennさんの「静的解析とコード生成」スライドの「型チェック」を読むと良いかも。)

## golangci-lintとの連携

https://golangci-lint.run/plugins/go-plugins

### [linterをgolangci-lintへ公式に追加する](https://golangci-lint.run/contributing/new-linters/#how-to-add-a-public-linter-to-golangci-lint)
### [私的なlinterをgolangci-lintから呼び出す](https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint)

[Go Plugin System](https://golangci-lint.run/plugins/go-plugins)による方法はおすすめしない。理由は、[こちら](https://speakerdeck.com/kuro_kurorrr/golangci-lint-module-plugin-system)とかに書いてあるが、やってみるとわかるが、めちゃんこ面倒くさい。

[Module Plugin System](https://golangci-lint.run/plugins/module-plugins)がおすすめ。
とはいえ、この方法も「カスタムなgolangci-lintをビルドする」という部分が辛い(ビルドは30秒ぐらいかかるので、私的なLinterを書き換えたらビルドが必要というのは、けっこうダルい)ことは変わらないため、私的なlinterをgolangci-lintから実行しないようにする、という結論も普通にアリ。

## 教材

自分で作るよりも、すでにあるものを読むほうが早いかもれん。
go/analysisパッケージを用いて書かれたLinterの例

- [golang.org/x/tools/go/analysis/passes 配下にあるLinter](https://cs.opensource.google/go/x/tools/+/refs/tags/v0.26.0:go/analysis/passes/)
- [golangci-lintで提供されているLinterたち](https://github.com/golangci/golangci-lint/tree/master/pkg/golinters)

# Memo

## linterのスケルトンコード生成

```bash
NAME=linter01 make linter
```
