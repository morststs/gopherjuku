# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

# CLAUDE.md — Gopherjuku（Go 学習アプリ）

## プロジェクト概要

Go 言語を「読む・動かす・理解する」ための Windows デスクトップ学習アプリ（アプリ名: **Gopherjuku**）。
Wails v2 + Svelte 5 製の 3 ペイン IDE 風 UI で、エディタに表示された Go サンプルコードを
その場でビルド・実行し、結果を確認しながら学習する。

> 名称の注意: 「Gopher」は Go コミュニティで広く使われる一般語だが、Gopher マスコット画像を
> 使う場合は CC BY 3.0 の帰属表示が必要、かつ Google/Go 公式を装う表現は避ける
> （名前のみの使用なら該当しない）。商用配布時は JPO/USPTO の商標確認を推奨。

- **左ペイン:** レッスンツリー（カテゴリー → コードの題）。**カテゴリーはアコーディオンで開閉可**。
  カテゴリーは **基本 / 標準ライブラリ / 応用** の3つ（`_contents/` の連番ディレクトリ）。
  計 29 レッスン実装済み（基本12・標準9・応用8）。全レッスン yaegi 実行確認済み。
- **中央ペイン:** Monaco エディタ。選択したレッスンの Go ソースを表示・編集する。
  **Shift+Alt+F** で gofmt 整形（`App.Format` = `go/format` を呼ぶ整形プロバイダを登録済み）。
  **Ctrl+マウスホイール**で文字サイズを拡大縮小（Monaco の `mouseWheelZoom`）。
- **右ペイン:** 実行結果（標準出力・エラー）。**Ctrl+マウスホイール**で出力の文字サイズを変更可。
- **学習方式:** 説明文は独立したドキュメントではなく **ソースコードのコメントに記載** する。
  「コードを読んで、実行して、挙動で理解する」形式。
- **Go の実行:** **yaegi**（純 Go 製の Go インタプリタ）で実行する。yaegi はライブラリとして
  本体 exe に static リンクされるため **外部の Go ツールチェーン不要**。実行は「自分自身を
  子プロセス（インタプリタモード）として起動し、標準入力でソースを渡す」方式で、
  無限ループはタイムアウトで確実に kill できる。標準出力を右ペインへ返す。

> 環境・ビルド構成は `git@github.com:morststs/sirusita.git`（Wails + Svelte 5 メモアプリ）を
> 下敷きにしている。異なるのは主に **用途（学習アプリ）** と **yaegi による Go 実行系**。

> **実行系の方針変更（重要）:** 当初 input01.md では「GopherJS を使う」と指定されていたが、
> GopherJS は実行時に Go ツールチェーン（Go 1.21 + gopherjs）を必要とし **単体 exe で配布できない**
> ことが判明したため、ユーザー合意のうえ **yaegi（本体同梱のインタプリタ）** に変更した。
> これにより単体 exe だけで実行が完結する。

> フェーズ1実装済み: 3 ペイン UI・yaegi 実行パイプライン（子プロセス＋タイムアウト）・
> 「基本 > ハローワールド」1件。教材コンテンツの拡充は次フェーズ。

## 技術スタック

- **バックエンド:** Go 1.23 + Wails v2.12
- **フロントエンド:** Svelte 5（runes）+ Vite 7
- **UIフレームワーク:** Flowbite Svelte 1.x + TailwindCSS 4（`@tailwindcss/vite`）
- **エディタ:** `monaco-editor`（Go シンタックスハイライト・スリム構成）
- **Go 実行系:** yaegi（`github.com/traefik/yaegi`・純 Go インタプリタ・本体に同梱）
- **ライセンス:** MIT

## 開発環境

### Podman 必須（コンテナ内で完結）

ホストに Go/Node/Wails はインストールしない。ビルド・テストはすべて Podman コンテナ内で
実行する。Windows 用 exe のクロスビルドまでコンテナ内で完結する。

```bash
# Podman イメージのビルド（初回 or Dockerfile 変更時）
podman build -t wails-dev .

# Go コマンド実行
podman run --rm -v "$PWD":/app:Z -w /app wails-dev go test -v ./...

# Wails ビルド（Linux / 動作確認用）
podman run --rm -v "$PWD":/app:Z -w /app wails-dev wails build

# Wails ビルド（Windows / 配布物）
podman run --rm -v "$PWD":/app:Z -w /app wails-dev wails build -platform windows/amd64

# フロントエンド npm コマンド
podman run --rm -v "$PWD":/app:Z -w /app/frontend wails-dev npm install <package>
```

> `:Z` は SELinux 環境でのマウントラベル付け。不要な環境では外してよい。
> rootless Podman ではホストとコンテナの UID が揃うため、生成物への `chown` は基本不要。

### イメージ内容（Dockerfile / wails-dev）

sirusita の `Dockerfile` とほぼ同じ（追加ツールチェーンは不要）:

- Go 1.23, Node.js 22 LTS（NodeSource）, Wails CLI v2.12.0
- `libgtk-3-dev`, `libwebkit2gtk-4.0/4.1-dev`（Linux ビルド用）
- `gcc-mingw-w64-x86-64`, `nsis`（Windows クロスコンパイル用）

> **Go 実行系は yaegi（`go.mod` の依存）:** Go の実行は yaegi を本体に static リンクして行うため、
> GopherJS や 2 つ目の Go SDK のような追加ツールチェーンは **一切不要**。これがビルド構成を
> 単純化し、かつ **単体 exe だけで実行が完結**（配布先に何もインストール不要）する理由。

> Svelte 5 / Vite 7 は Node 20.19+ / 22.12+ を要求するため、イメージは Node 22 を使用。

## プロジェクト構成

> Go モジュール名は `gopherjuku`（`go.mod`）、`wails.json` の `outputfilename` も `gopherjuku`。

```
gopherjuku/
├── CLAUDE.md                # 本ファイル
├── SKILLS.md                # 反復作業のプレイブック（レッスン追加手順など）
├── doc/input01.md           # 要件メモ（フェーズ入力）
├── main.go                  # Wails エントリポイント。os.Args[1]==__interp なら子プロセスとして
│                            #   yaegi 実行（RunChild）。それ以外は GUI 起動。frontend/dist と
│                            #   _contents を go:embed。Windows タイトルバーの色も指定（CustomTheme）
├── app.go                   # App 構造体（Tree/Source/Run/Format を Wails にバインド）
├── internal/lessons/
│   ├── lessons.go           # 教材の読み込み。_contents を走査し、数字プレフィックス順に
│   │                        #   カテゴリー/題を組み立てる。Wails 非依存で単体テスト可
│   └── lessons_test.go      # 並び順・表示名・パス・parseOrder のテスト（fstest.MapFS）
├── internal/runner/
│   ├── runner.go            # yaegi 実行。Run=子プロセス起動+タイムアウト、Interpret=同一プロセス、
│   │                        #   RunChild=子プロセス側入口。Wails 非依存で単体テスト可
│   └── runner_test.go       # yaegi 実行のテスト（in-process・子プロセス・タイムアウト）
├── internal/gofmt/
│   ├── gofmt.go             # go/format(gofmt) の薄いラッパー（Shift+Alt+F の整形で使用）
│   └── gofmt_test.go        # 整形/構文エラーのテスト
├── Dockerfile               # ビルド用イメージ（Podman・Go 1.23 + Node 22 + Wails。追加toolchain不要）
├── wails.json               # Wails 設定
├── _contents/               # レッスン教材。"_" 始まりで go ビルド対象外（go:embed で同梱）
│   ├── 1_基本/              #   ディレクトリ/ファイル先頭の連番で表示順を制御（表示名からは除去）
│   │   ├── 01 ハローワールド.go
│   │   ├── … 02 変数と型 / 05 制御構文 / 12 インターフェース …
│   ├── 2_標準ライブラリ/    #   02 strings / 04 sort / 08 JSON …
│   └── 3_応用/              #   03 エラーハンドリング / 05 goroutine と channel / 08 ジェネリクス …
├── frontend/
│   ├── svelte.config.js     # vitePreprocess({ script: true })（下記「ビルド注意点」）
│   ├── vite.config.js       # Vite + svelte + @tailwindcss/vite
│   ├── index.html
│   └── src/
│       ├── main.js          # Svelte 5 マウント
│       ├── app.css          # Tailwind 4 + flowbite @source + ベーススタイル
│       ├── App.svelte       # ルート（3ペインのスプリッター + 状態管理 + 実行フロー）
│       ├── LessonTree.svelte # 左: カテゴリー/題ツリー
│       ├── Editor.svelte     # 中央: Monaco エディタ（Go）
│       ├── Output.svelte     # 右: 実行結果
│       ├── wails.js          # バックエンド呼び出しラッパー（window.go 経由・ブラウザ単体フォールバック）
│       └── monaco.js         # Monaco スリム構成（エディタ + Worker + go 用 gofmt 整形プロバイダ）
├── LICENSE                  # MIT
└── THIRD_PARTY_LICENSES.md  # （教材拡充・依存追加時に整備）
```

## 実行フロー

1. ユーザーが左ペインでレッスンを選ぶ → 中央 Monaco に Go ソースを読み込む（`App.Source`）。
2. 「実行」で中央エディタの現在のソースを `App.Run`（Wails バインディング）へ渡す。
3. `internal/runner` の `Run` が **自分自身の実行ファイルを子プロセスとして起動**する
   （`os.Executable()` + 引数 `__interp`）。ソースは標準入力で渡す。
4. 子プロセス（`main.go` が `RunChild` へ分岐）が **yaegi** でソースを評価し、
   プログラムの標準出力を自身の stdout へ、コンパイル/実行エラーを stderr へ書く。
5. 親は stdout/stderr を集めて `Result{Success, Output, Error}` を返す。
   - **タイムアウト**（既定 5 秒）を超えたら `context` で子プロセスを kill →「無限ループの
     可能性」を返す。インプロセスの goroutine では強制停止できないため、この子プロセス方式を採る。
6. フロントは `Output`/`Error` を右ペインに表示するだけ（**JS の実行・iframe は無し**）。

## ブラウザ版（GitHub Pages）

同じフロントエンドを、バックエンド無しの静的サイトとしても配信する
（公開先: https://morststs.github.io/gopherjuku/ ・`gh-pages` ブランチ）。

- **実行/整形**: yaegi と `go/format` を **WebAssembly 化**（`web/wasm/`・js/wasm 専用の
  入れ子モジュール）し、`frontend/public/worker.js`（Web Worker）内で動かす。無限ループは
  ワーカーを terminate して停止（デスクトップの子プロセス kill に相当）。
- **レッスン**: `cmd/gen-lessons` が `_contents/` を静的 JSON（`lessons.json`）に書き出す。
- **切り替え**: `frontend/src/wails.js` が `window.go` の有無でデスクトップ/ブラウザを判定。
  ブラウザ時のみ `webRunner.js` を動的 import。
- **注意**: `main.wasm`/`wasm_exec.js`/`lessons.json` は生成物（`.gitignore` 済み）。
  デスクトップ exe に混入させないため、`wails build` 前に `frontend/public/` から消す
  （手順は SKILLS #9）。ビルドは Pages 用に `vite build --base=/gopherjuku/`。

## ビルド注意点（svelte.config.js）

`frontend/svelte.config.js` は **必須**。`vitePreprocess({ script: true })` を有効化する。
flowbite-svelte は TypeScript 入りの `.svelte` を配布しており、Svelte 5 の組み込み TS 除去では
一部の型注釈が残って Rollup ビルドが失敗する。esbuild による script トランスパイルを明示して
取り込めるようにする。**この設定を削除するとビルドが壊れる。**

## コーディング規約

- Go: 標準フォーマット（`gofmt`）。
- Svelte 5（runes）: props は `$props()`、状態は `$state`/`$derived`/`$effect`、
  親への通知はコールバック props（`onXxx`）、イベントは `onclick` 等のネイティブ属性。
- **教材（`_contents/**/*.go`）のコメントは日本語**。学習者が読む説明はコードコメントに書く。
- コード内コメントは日本語 OK、**識別子は英語**。
- コミットメッセージ: `feat:` / `fix:` / `chore:` プレフィックス。
- Go API を変更したら Wails バインディングを再生成:
  `podman run --rm -v "$PWD":/app:Z -w /app wails-dev wails generate module`

## セキュリティ / 隔離

- ユーザー入力の Go ソースを実行するため、**信頼境界に注意**。実行は **子プロセス**に隔離し、
  **タイムアウト**（既定 5 秒・`internal/runner` の `defaultTimeout`）で無限ループを kill する。
- 現状 yaegi には標準ライブラリ全体（`stdlib.Symbols`）を公開している。ユーザーは自分の PC で
  自分のコードを動かす前提（`go run` 相当）だが、より厳しくしたい場合は公開シンボルを
  `fmt`/`strings`/`math` 等に絞れば `os/exec`・`net` 等を封じられる（`runner.interpret` を参照）。
- yaegi はインタプリタのため **Go の完全なサブセット**。基本文法・goroutine・interface・
  エラーラップ・reflect・主要 stdlib（json/sort/strings/strconv/time/errors 等）は動作確認済み。
  **既知の制限:**
  - **標準入力は読めない**（子プロセスの stdin をソース供給に使うため）。`os.Stdin`/`fmt.Scan`
    系の教材は不可。入力はハードコードで示す。
  - **ジェネリクスの型推論が不完全**：戻り値型が別の型パラメータになる推論
    （例 `Map[T,U]` で int→string を推論）は失敗する。**明示的型引数**
    （`Map[int,string](...)`）や、単一型パラメータ（`Filter[T]`）・制約付き（`Sum[T Number]`）
    の形にすれば動く。ジェネリック型（`Stack[T]`）は可。
  - **`errors.As` は独自エラー型で panic する**（`*target must be interface or implement error`）。
    `errors.Is`＋sentinel（`%w` ラップ）は動くので、そちらを使う。
  - **`slices` パッケージはほぼ未対応**（`Contains`/`Index` のみ可。`Sort`/`Max`/`Reverse`/
    `DeleteFunc` 等は不可。組み込み `clear` も無い）。`maps` も同様とみて避ける。
    並べ替えは `sort` パッケージを使う。
  - 教材を追加したら **必ず実機バイナリで実行確認**すること
    （`cat "<file>" | ./build/bin/<bin> __interp`）。実測で 29 レッスンすべて通過を確認済み。
- **配布は単体 exe で完結**（yaegi 同梱）。GopherJS 方式のような実行時ツールチェーン依存は無い。

## ライセンス

- **MIT**（`LICENSE`）。サードパーティは `THIRD_PARTY_LICENSES.md` を参照。
- 秘密鍵・署名資材（`certs/` 等）はコミットしない（`.gitignore`）。
