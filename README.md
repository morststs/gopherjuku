<img src="build/appicon.png" width="96" alt="Gopherjuku">

# Gopherjuku — Go を動かして学ぶ

Go 言語のサンプルコードを **その場で編集して実行し、結果を見ながら学ぶ** 学習アプリです。
説明はソースコードのコメントに書かれており、「読む → 動かす → 理解する」形式で進めます。

## 使い方は 2 通り

### 🌐 ブラウザ版（インストール不要）

**https://morststs.github.io/gopherjuku/**

GitHub Pages 上でそのまま動きます。Go の実行は **yaegi インタプリタを WebAssembly 化**して
ブラウザ内で行うため、サーバーもインストールも不要です。
（初回のみ実行環境の wasm を読み込みます。少し大きめのダウンロードがあります。）

### 🖥 デスクトップ版（Windows）

単体の `.exe` として動きます。実行エンジン（yaegi）を同梱しているので、
**別途 Go のインストールは不要**です。ビルド方法は下記「開発」を参照してください。

## 画面

3 ペイン構成です。

| ペイン | 内容 |
|--------|------|
| 左 | レッスンツリー（カテゴリーはアコーディオンで開閉） |
| 中央 | Monaco エディタ（Go 構文ハイライト） |
| 右 | 実行結果（標準出力・エラー） |

- **▶ 実行** … コードを実行して結果を右ペインに表示
- **Shift + Alt + F** … gofmt で整形
- **Ctrl + マウスホイール** … エディタ／実行結果の文字サイズ変更

## 収録レッスン（計 29）

- **基本** … ハローワールド / 変数と型 / 定数と iota / 演算子 / 制御構文 / 関数 /
  クロージャ / 配列とスライス / マップ / 構造体 / メソッド / インターフェース
- **標準ライブラリ** … fmt / strings / strconv / sort / time / math と乱数 /
  errors / JSON / regexp
- **応用** … 型アサーションと型スイッチ / 埋め込み / エラーハンドリング /
  defer・panic・recover / goroutine と channel / select / sync / ジェネリクス

## 技術構成

| | |
|---|---|
| バックエンド（デスクトップ） | Go 1.23 + Wails v2 |
| フロントエンド | Svelte 5（runes）+ Vite 7 + TailwindCSS 4 |
| エディタ | Monaco Editor |
| Go 実行 | **yaegi**（純 Go 製インタプリタ）— デスクトップは子プロセス、ブラウザは WebAssembly |
| 整形 | `go/format`(gofmt) |

実行エンジンに yaegi を採用したことで、外部の Go ツールチェーンなしに
**単体 exe / 静的サイトだけで実行が完結**します。

## 開発

ビルドはすべて Podman コンテナ内で完結します（ホストに Go/Node は不要）。

```bash
# ビルド用イメージ
podman build -t wails-dev .

# テスト
podman run --rm -v "$PWD":/app:Z -w /app wails-dev go test ./...

# Windows 版 exe（build/bin/gopherjuku.exe）
podman run --rm --network=host -v "$PWD":/app:Z -w /app wails-dev \
  wails build -platform windows/amd64
```

ブラウザ版（GitHub Pages）のビルド・レッスン追加・アイコン再生成などの手順は
[`SKILLS.md`](SKILLS.md)、設計や yaegi の既知の制限は [`CLAUDE.md`](CLAUDE.md) を参照してください。

## ライセンス

[MIT](LICENSE)
