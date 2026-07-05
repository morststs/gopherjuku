// fmt — 書式付き入出力
//
// Println / Printf / Sprintf と、代表的な「動詞（verb）」を確認します。
package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	// Println：スペース区切りで出力して改行。
	fmt.Println("値を並べる:", 1, "two", 3.0, true)

	p := Point{X: 1, Y: 2}

	// 代表的な verb
	fmt.Printf("%%v  値      : %v\n", p)
	fmt.Printf("%%+v フィールド名付き: %+v\n", p)
	fmt.Printf("%%d 整数 / %%b 2進 / %%x 16進: %d %b %x\n", 255, 255, 255)
	fmt.Printf("%%s 文字列 / %%q 引用符付き: %s %q\n", "go", "go")
	fmt.Printf("%%f 小数 / %%.2f 桁指定: %f %.2f\n", 3.14159, 3.14159)
	fmt.Printf("%%T 型: %T %T\n", 42, "hi")

	// Sprintf：出力せず文字列として組み立てる。
	msg := fmt.Sprintf("(%d, %d)", p.X, p.Y)
	fmt.Println("Sprintf の結果:", msg)
}
