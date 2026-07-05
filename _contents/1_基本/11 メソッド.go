// メソッド
//
// メソッドは特定の型に結び付いた関数です。レシーバに「値」を使うか
// 「ポインタ」を使うかで、元の値を変更できるかが変わります。
package main

import "fmt"

type Counter struct {
	count int
}

// 値レシーバ：コピーに対して働くので、元の値は変わらない（読み取り向き）。
func (c Counter) Value() int {
	return c.count
}

// ポインタレシーバ：元の値を変更できる（更新向き）。
func (c *Counter) Increment() {
	c.count++
}

func main() {
	c := Counter{}

	// ポインタレシーバのメソッドは、変数に対してそのまま呼べる（自動で &c になる）。
	c.Increment()
	c.Increment()
	c.Increment()

	fmt.Println("カウント:", c.Value()) // 3

	// 値レシーバのメソッドは元を変えない例。
	c2 := c
	c2.Increment()
	fmt.Println("c:", c.Value(), "c2:", c2.Value()) // 3 4
}
