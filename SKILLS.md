# SKILLS.md — 反復作業プレイブック（Gopherjuku）

このリポジトリで繰り返し発生する作業の手順集。各「スキル」は独立した手順として実行できる。
環境はすべて Podman コンテナ（`wails-dev`）内で完結する（`CLAUDE.md` の「開発環境」参照）。

> フェーズ1実装済み（3ペインUI・yaegi 実行・「基本>ハローワールド」）。★印は「アプリの実体を
> 触る手順」の目印（フェーズ0の名残。現在はいずれも実行可能）。

---

## スキル一覧

| # | スキル | いつ使うか |
|---|--------|-----------|
| 1 | 開発環境を用意する | 初回 / `Dockerfile` を変更したとき |
| 2 | ★レッスンを追加する | 新しいコードの題（サンプル）を増やすとき |
| 3 | ★カテゴリーを追加する | 左ペインに新しい分類を足すとき |
| 4 | ★アプリを実行して動作確認 | UI/挙動を確認するとき |
| 5 | ★Windows exe をビルド | 配布物を作るとき |
| 6 | ★Wails バインディング再生成 | Go の公開メソッドを変えたとき |
| 7 | テストを実行する | 変更後の確認 |
| 8 | アイコンを作り直す | ロゴ/アイコンを変えるとき |

---

## 1. 開発環境を用意する

```bash
# ビルド用イメージを作成
podman build -t wails-dev .
```

`Dockerfile` は sirusita とほぼ同じ（Go 1.23 + Node 22 + Wails CLI）。Go の実行は yaegi を
本体に同梱するため追加ツールチェーンは不要。イメージが古いと感じたら再ビルドする。

> podman-in-container 環境ではビルド/実行時に `--isolation=chroot`（build）や `--network=host`
> が必要になることがある（`podman build --isolation=chroot -t wails-dev .` 等）。

---

## 2. ★レッスンを追加する

教材は `_contents/<連番_カテゴリー>/<連番 題>.go` として置く。**説明は独立ドキュメントにせず、
ソースのコメントに日本語で書く**（本アプリの学習方式）。

> ディレクトリ名が `_contents`（先頭 `_`）なのは、埋め込む `.go` を Go のビルド対象として
> 走査させないため（go ツールは `_` 始まりのディレクトリを無視、`go:embed all:` は取り込む）。

**表示順は先頭の連番で決まる**（表示名からは連番を除去して表示）。例:

```
_contents/
├── 1_基本/
│   ├── 01 ハローワールド.go
│   ├── 02 変数と型.go
│   └── 12 インターフェース.go   ← 「12」の位置に並ぶ。表示は「インターフェース」
├── 2_標準ライブラリ/
└── 3_応用/
```

手順:
1. `_contents/<連番_カテゴリー>/<NN 題>.go` を作成し、コメント付きで動く Go サンプルを書く。
   連番は既存と重複しないよう空きを空けて付けてよい（間に挿入しやすい）。
2. **yaegi（インタプリタ）で動く範囲に留める**。主要機能は動くが既知の制限あり:
   - **標準入力は読めない**（`os.Stdin`/`fmt.Scan` 系は不可。入力はハードコードで示す）
   - **ジェネリクスの型推論が不完全**：戻り値が別の型パラメータになる推論は失敗。
     明示的型引数（`Map[int,string](...)`）や単一型パラメータ（`Filter[T]`）にする。
   - `unsafe`・cgo は不可。
3. 再ビルドすると `main.go` の `go:embed all:_contents` が自動で取り込む。一覧登録は不要
   （`internal/lessons` が走査し、カテゴリー=ディレクトリ・題=`.go` 名・順序=連番で組み立てる）。
4. **必ず動作確認する**。ビルド済みバイナリに直接流すのが手早い:
   `cat "_contents/1_基本/01 ハローワールド.go" | ./build/bin/<バイナリ> __interp`

---

## 3. ★カテゴリーを追加する

`_contents/` 直下に **連番付き**ディレクトリを作る（例: `_contents/3_応用/`）。
左ペインのツリーはカテゴリー = ディレクトリ、題 = `.go` ファイル名、順序 = 連番で構成される。
表示名は連番プレフィックスを除いた部分（`3_応用` → 「応用」）。

---

## 4. ★アプリを実行して動作確認

```bash
# Linux 上で Wails 開発起動（GUI 確認）
podman run --rm -v "$PWD":/app:Z -w /app wails-dev wails dev
```

確認観点: 左ペインに「基本 > ハローワールド」が出る / 中央 Monaco にソースが表示される /
「実行」で右ペインに標準出力が出る / コンパイルエラーが右ペインに表示される。

---

## 5. ★Windows exe をビルド

```bash
podman run --rm -v "$PWD":/app:Z -w /app wails-dev wails build -platform windows/amd64
# 出力: build/bin/gopherjuku.exe
```

---

## 6. ★Wails バインディング再生成

Go の公開メソッド（`app.go` の `Tree` / `Source` / `Run` など）を変更したら:

```bash
podman run --rm -v "$PWD":/app:Z -w /app wails-dev wails generate module
```

`frontend/wailsjs/` は自動生成物。手で編集しない（※本アプリは `window.go` 経由で呼ぶため
`wailsjs` の import には依存しないが、型が欲しい場合はこれで生成できる）。

---

## 7. テストを実行する

```bash
podman run --rm -v "$PWD":/app:Z -w /app wails-dev go test -v ./...
```

`internal/runner` のテストは yaegi の実行（同一プロセス・子プロセス・タイムアウト kill）を
実際に検証する。外部ツールチェーン不要。

---

## 8. アイコンを作り直す

アプリアイコンの元画像は `build/appicon.png`。生成プログラムが `build/icongen/`（親モジュールの
ビルド対象外の入れ子モジュール）にあるので、デザインを変えたら再生成する:

```bash
# 元画像を再生成
podman run --rm -v "$PWD":/app:Z -w /app/build/icongen wails-dev go run . /app/build/appicon.png
# 既存の .ico を消して、ビルド時に新 appicon から作り直させる
rm -f build/windows/icon.ico
# 再ビルド（wails が appicon.png から icon.ico を生成し、exe に埋め込む）
podman run --rm --network=host -v "$PWD":/app:Z -w /app wails-dev wails build -platform windows/amd64
```

> アプリ内トップバーのロゴは `frontend/src/App.svelte` のインライン SVG。アイコンを変えたら
> こちらも合わせると見た目が揃う。
