// Package gofmt は Go ソースを標準の gofmt 整形にかける薄いラッパー。
// go/format（標準ライブラリ）を使うだけなので外部ツール不要・Wails 非依存で、
// 単体テストできる。エディタの「Format Document」(Shift+Alt+F) から呼ばれる。
package gofmt

import "go/format"

// Result は整形結果。フロントエンド（Wails バインディング）へ JSON で渡される。
type Result struct {
	Success   bool   `json:"success"`   // 整形できたら true
	Formatted string `json:"formatted"` // 整形後のソース（失敗時は入力をそのまま返す）
	Error     string `json:"error"`     // 整形失敗（構文エラー等）のメッセージ
}

// Source は src を gofmt 整形する。構文エラー等で整形できない場合は
// Success=false とし、Formatted には入力をそのまま入れて返す（呼び出し側で
// 「変更しない」判断ができるように）。
func Source(src string) Result {
	out, err := format.Source([]byte(src))
	if err != nil {
		return Result{Success: false, Formatted: src, Error: err.Error()}
	}
	return Result{Success: true, Formatted: string(out)}
}
