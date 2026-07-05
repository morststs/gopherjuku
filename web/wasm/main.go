//go:build js && wasm

// ブラウザ（GitHub Pages）用の Go ランナー。
// yaegi インタプリタと go/format(gofmt) を WebAssembly にコンパイルし、
// JS のグローバル関数として公開する。デスクトップ版（Wails）の
// バックエンド（internal/runner・internal/gofmt）と同じ役割をブラウザ内で担う。
//
// このモジュールは js/wasm 専用（build タグ）なので、通常の `go build ./...`
// の対象外（親モジュールからも独立した入れ子モジュール）。
package main

import (
	"bytes"
	goformat "go/format"
	"syscall/js"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// gopherjukuRun(source) -> { output, error }
// yaegi でソースを実行し、標準出力（+標準エラー）とエラーを返す。
func runGo(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return js.ValueOf(map[string]any{"output": "", "error": "no source"})
	}
	var buf bytes.Buffer
	i := interp.New(interp.Options{Stdout: &buf, Stderr: &buf})
	if err := i.Use(stdlib.Symbols); err != nil {
		return js.ValueOf(map[string]any{"output": "", "error": err.Error()})
	}
	res := map[string]any{"output": "", "error": ""}
	if _, err := i.Eval(args[0].String()); err != nil {
		res["error"] = err.Error()
	}
	res["output"] = buf.String()
	return js.ValueOf(res)
}

// gopherjukuFormat(source) -> { formatted, ok }
// go/format(gofmt) でソースを整形する。整形できなければ ok:false。
func formatGo(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return js.ValueOf(map[string]any{"formatted": "", "ok": false})
	}
	src := args[0].String()
	out, err := goformat.Source([]byte(src))
	if err != nil {
		return js.ValueOf(map[string]any{"formatted": src, "ok": false})
	}
	return js.ValueOf(map[string]any{"formatted": string(out), "ok": true})
}

func main() {
	js.Global().Set("gopherjukuRun", js.FuncOf(runGo))
	js.Global().Set("gopherjukuFormat", js.FuncOf(formatGo))
	select {} // 常駐して JS からの呼び出しに応える
}
