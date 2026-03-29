---
paths:
  - "**/*.go"
---

# Coding Conventions

- **Formatting**: `gofmt` を使用する。コミット前に必ず `gofmt -w .` を実行すること。
- **Testing**: テーブルドリブンテスト（`[]struct` + `t.Run`）を使用する。`greet_test.go` をリファレンスとする。
- **Naming**: エクスポートされる関数には英語の GoDoc コメントをつける。
- **Error handling**: エラーは `errors.New()` または `fmt.Errorf()` で生成する。
- **Git branching**: `feat/<issue-number>-<description>` または `fix/<issue-number>-<description>`。
