---
paths:
  - "**/*.go"
---

# Architecture

## Package Structure

- **Root package (`sandbox`)**: 単一パッケージのフラット構成。サブパッケージはない。
- 各機能は独立したファイルに分離する（例: `greet.go`, `math.go`）
- テストは対応する `*_test.go` に配置する

## Design Patterns

- 引数が空文字の場合はデフォルト値を使用するパターン（`greet.go` 参照）
- エラーが起こりうる関数は `(T, error)` を返す（`math.go` の `Divide` 参照）
