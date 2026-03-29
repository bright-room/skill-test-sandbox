# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

dev-workflow スキル（implement-plan, fix-review, review 等）のE2Eテスト用サンドボックス。最小限の Go パッケージで構成される。

## Development Commands

```bash
# ビルド
go build ./...

# テスト
go test ./...

# フォーマット
gofmt -w .
```

## Architecture

単一パッケージ (`package sandbox`) のフラットな構成。モジュールやサブパッケージはない。

## Coding Guidelines

- 既存の関数（`Greet`, `Farewell` 等）と同じパターンに従う
- テストはテーブルドリブンテスト（`[]struct` + `t.Run`）で書く
- エラーを返す関数は `(T, error)` のシグネチャにする
