// ブラウザ（GitHub Pages）用の WebAssembly ランナー専用モジュール。
// 独自 go.mod を持つ入れ子モジュールにして、親（gopherjuku）の
// `go build ./...` / `go vet ./...` の対象外にする（js/wasm 専用のため）。
module gopherjuku/web/wasm

go 1.23.0

require github.com/traefik/yaegi v0.16.1
