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

## 参考
- 公式: https://golangci-lint.run/
- 公式Github: https://github.com/golangci/golangci-lint
- https://qiita.com/twrcd1227/items/9d62a5ac58232d99b656