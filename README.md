# Anemone: A common command framework

Anemone is a code base for the lambda function working on the clamshell

- RESTfullとしてのWebからのリクエスト処理をCLIからも同様に受付・処理できる
- 上記によりWebAPI開発におけるデバッグなどの煩わしさに関する問題を接続I/Fとロジック実装の二つのテーマに分割する
- 本フレームワークでは接続I/Fを透過的に支援する機構とロジック追加を容易にする機構をフレームワークとし提供する事にある

# DEMO

# Features

# Requirement

## Development Reuqirements

* cobra 

```
# go get -u github.com/spf13/cobra/cobra
※ cobraのPATHが通ってる必要があります。
  $GOPATH/bin/cobra
```

## How to preview godoc
godoc -http=:8080 -goroot=/path/to/yourdir/anemone

# Installation


# Usage

## How to start Development

