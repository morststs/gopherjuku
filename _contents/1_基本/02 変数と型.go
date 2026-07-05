// 変数と型
//
// Go の変数宣言と基本型を見ていきます。値を書き換えて実行し、
// 出力がどう変わるか試してみてください。
package main

import "fmt"

func main() {
	// var で宣言。初期値を省くと「ゼロ値」（int は 0、string は ""）になる。
	var age int
	var name string = "Gopher"
	fmt.Println("age(ゼロ値):", age, "/ name:", name)

	// := は型推論付きの短い宣言（関数の中でだけ使える）。
	height := 172.5 // float64 と推論される
	active := true  // bool
	fmt.Printf("height=%.1f active=%v\n", height, active)

	// 主な基本型
	var i int = -42
	var f float64 = 3.14
	var s string = "文字列"
	var r rune = 'あ' // rune は Unicode コードポイント（int32 の別名）
	var b byte = 'A'  // byte は uint8 の別名
	fmt.Println("int/float/string:", i, f, s)
	fmt.Printf("rune %c=%d, byte %c=%d\n", r, r, b, b)

	// 型変換は必ず明示的に書く（暗黙の変換は無い）。
	n := 10
	x := float64(n) / 3
	fmt.Printf("%d / 3 = %.3f\n", n, x)
}
