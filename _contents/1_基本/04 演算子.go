// 演算子
//
// 算術・比較・論理・ビット演算の基本を確認します。
package main

import "fmt"

func main() {
	a, b := 17, 5

	// 算術演算：/ は整数どうしだと整数除算、% は剰余。
	fmt.Println("和 差 積:", a+b, a-b, a*b)
	fmt.Println("商 余り:", a/b, a%b)

	// 比較演算：結果は bool。
	fmt.Println("a>b, a==b:", a > b, a == b)

	// 論理演算：&&（かつ）, ||（または）, !（否定）。
	t, f := true, false
	fmt.Println("&& || !:", t && f, t || f, !t)

	// ビット演算：& | ^ << >>
	x, y := 0b1100, 0b1010
	fmt.Printf("AND=%04b OR=%04b XOR=%04b\n", x&y, x|y, x^y)
	fmt.Printf("左シフト 1<<4 = %d\n", 1<<4)
}
