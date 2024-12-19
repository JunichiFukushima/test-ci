## 概要

静的解析の検証用repo
目的や背景については以下参照
[目次(Docleプロジェクトの静的解析導入方針)](https://schooinc.atlassian.net/wiki/spaces/ServiceDevelopment/pages/3375530243/Docle)


## 方針(現状)

- golangci-lintを使う
    - .golangci.ymlに設定ファイルを置いておく
    - vsコード / GoLandに全員共通の拡張機能を入れておくようにする


## ローカルで動かす手順

golangci-lint のインストール
```
brew install golangci-lint
```

※go installでインストールするのは非推奨


実行

```
golangci-lint run
```

## vsコードの拡張機能設定手順

- settingsに以下の設定を入れる

```
{
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--config=${workspaceFolder}/.golangci.toml", "--fast"]
}
```


## FYI

- デバッグ付きで実行
```
golangci-lint run -v
```

- 設定可能な、lintersを出力
```
golangci-lint linters
```

## 設定可能な内容


```
asasalint

内容: 可変長関数の引数で []any を ...any に変換すべき箇所を検出します。
asciicheck

内容: コード内の識別子が非ASCII文字を含んでいないかをチェックします。
bidichk

内容: 危険なUnicode文字列シーケンスを検出します（双方向テキストの制御文字など）。
bodyclose

内容: HTTPレスポンスのボディが正しく閉じられているか確認します。
canonicalheader

内容: net/http.Header が正しい標準ヘッダー形式を使用しているかをチェックします。
containedctx

内容: 構造体内に context.Context フィールドが含まれている箇所を検出します。
contextcheck

内容: 非継承の context.Context を使用している箇所をチェックします。
copyloopvar

内容: ループ変数をコピーしている箇所を検出します。
cyclop

内容: 関数やパッケージの循環的複雑度を測定し、一定基準を超えた場合に警告します。
decorder

内容: 型、定数、変数、関数の宣言順序や数をチェックします。
depguard

内容: インポートされたパッケージが許可リストまたは禁止リストに含まれているかをチェックします。
dogsled

内容: 代入時に使われる空白識別子（例: _）が多すぎる箇所を検出します。
dupl

内容: 重複したコード（コードクローン）を検出します。
dupword

内容: ソースコード内の重複した単語を検出します。
durationcheck

内容: 2つの期間値を掛け合わせている箇所を検出します。
err113

内容: Go 1.13以降でのエラーラップに関する問題を検出します。
errchkjson

内容: JSONエンコード関数に渡される型のチェックや、エラーの未処理箇所を検出します。
errname

内容: エラー名が「Err」で始まるか、エラー型が「Error」で終わるかをチェックします。
errorlint

内容: Go 1.13以降でのエラーラッピングの問題を検出します。
exhaustive

内容: 列挙型のswitch文で全てのケースが網羅されているかを検出します。
exhaustruct

内容: 構造体の全フィールドが初期化されているかをチェックします。
fatcontext

内容: ループや関数内でネストされた context を検出します。
forbidigo

内容: 使用禁止の識別子を検出します。
forcetypeassert

内容: 強制的な型アサーションを使用している箇所を検出します。
funlen

内容: 長すぎる関数を検出します。
gci

内容: Goのインポート順序を強制し、一貫性を保ちます。
ginkgolinter

内容: ginkgoとgomegaの使用基準を適用します。
gocheckcompilerdirectives

内容: Goコンパイラディレクティブ（例: //go:）が有効かをチェックします。
gochecknoglobals

内容: グローバル変数が存在しないことを確認します。
gochecknoinits

内容: init関数が存在しないことを確認します。
gocognit

内容: 関数の認知的複雑度を計算し、一定基準を超えた場合に警告します。
goconst

内容: 繰り返し使われる文字列や数値を定数に置き換えるべき箇所を検出します。
godot

内容: コメントの末尾がピリオドで終わっているかをチェックします。
gofumpt

内容: gofumpt に基づくフォーマットを適用します。
goimports

内容: インポート文が goimport コマンドに従っているかを確認します。
gomoddirectives

内容: go.mod 内の replace、retract、exclude 指令を管理します。
gosimple

内容: より単純化できるコードを検出します。
ineffassign

内容: 未使用の変数代入を検出します。
misspell

内容: 英語の一般的なスペルミスを検出します。
nestif

内容: ネストが深いif文を検出します。
prealloc

内容: スライスを事前割り当てできる箇所を検出します。
wrapcheck

内容: 外部パッケージから返されるエラーが適切にラップされているかを確認します。
```

## 参考
- 公式: https://golangci-lint.run/
- 使い方: https://golangci-lint.run/usage/linters/
- 公式Github: https://github.com/golangci/golangci-lint
- https://qiita.com/twrcd1227/items/9d62a5ac58232d99b656