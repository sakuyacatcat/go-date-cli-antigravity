# Go Date CLI

`date` コマンドライクな Go 製 CLI ツールです。
現在時刻の表示に加え、タイムゾーンの変換や出力フォーマットのカスタマイズをサポートしています。

## 機能

- **現在時刻の表示**: デフォルトでシステムの現在時刻を表示します。
- **タイムゾーン変換**: `-z` / `--timezone` フラグで任意のタイムゾーン（例: `UTC`, `Asia/Tokyo`）を指定可能。
- **フォーマット指定**: `-f` / `--format` フラグで出力形式を指定可能。
  - カスタムレイアウト（例: `2006-01-02`）
  - プリセット: `ISO8601` (例: `2026-01-04T12:00:00Z`)

## 必要要件

- Go 1.20+
- Make (ビルドおよびテスト実行用)

## ビルド

```bash
make build
# => date-cli.exe が生成されます
```

## テスト

```bash
make test
# => 全ユニットテストを実行します
```

## 使い方

### 基本的な使用法

```bash
./date-cli.exe
# => Sun Jan  4 22:00:00 JST 2026
```

### タイムゾーンを指定する

```bash
./date-cli.exe -z "UTC"
# => Sun Jan  4 13:00:00 UTC 2026

./date-cli.exe --timezone "Asia/Tokyo"
# => Sun Jan  4 22:00:00 JST 2026
```

### フォーマットを指定する

```bash
# ISO8601 形式
./date-cli.exe -f "ISO8601"
# => 2026-01-04T22:00:00+09:00

# カスタムレイアウト (Go reference time: Mon Jan 2 15:04:05 MST 2006)
./date-cli.exe --format "2006/01/02 15:04:05"
# => 2026/01/04 22:00:00
```

### 組み合わせ

```bash
./date-cli.exe -z "UTC" -f "ISO8601"
# => 2026-01-04T13:00:00Z
```

## アーキテクチャ

このプロジェクトは、テスト容易性と保守性を考慮したレイヤードアーキテクチャを採用しています。

- **Layer Structure**:
  - `internal/infrastructure/clock`: `time` パッケージへの依存を抽象化した `Clock` インターフェースを提供。
  - `internal/service`: 時刻取得とゾーン変換を行うビジネスロジック。`Clock` インターフェースに依存。
  - `internal/view`: 出力フォーマットの整形ロジック。
  - `main.go`: CLI エントリーポイント。Dependency Injection (DI) により各コンポーネントを組み立てます。

- **Testing**:
  - ユニットテストでは `MockClock` や `MockTimeService` を使用し、システム時刻や OS に依存しない堅牢なテストを実現しています。
