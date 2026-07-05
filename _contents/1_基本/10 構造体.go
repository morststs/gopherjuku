// 構造体
//
// 構造体は複数の値をひとまとめにした型です。
package main

import "fmt"

type Point struct {
	X, Y int
}

type Person struct {
	Name string
	Age  int
	Home Point // 構造体は入れ子にできる
}

func main() {
	// フィールド名を指定した初期化（推奨）。
	p := Person{
		Name: "Alice",
		Age:  30,
		Home: Point{X: 1, Y: 2},
	}

	// フィールドへのアクセス・更新。
	p.Age++
	fmt.Println("名前:", p.Name, "年齢:", p.Age)
	fmt.Println("家の座標:", p.Home.X, p.Home.Y)

	// %+v はフィールド名付きで構造体を出力する。
	fmt.Printf("%+v\n", p)

	// ポインタ経由で更新すると元の値が変わる。
	move(&p.Home)
	fmt.Println("移動後:", p.Home)
}

func move(pt *Point) {
	pt.X += 10
	pt.Y += 10
}
