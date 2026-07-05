// インターフェース
//
// インターフェースは「メソッドの集合」を表す型です。ある型がそのメソッドを
// 実装していれば、明示的な宣言なしに自動でインターフェースを満たします
//（構造的部分型）。これが Go の多態性（ポリモーフィズム）の基本です。
package main

import "fmt"

// Shape は「Name() と Area() を持つ」という約束を表すインターフェース。
type Shape interface {
	Name() string
	Area() float64
}

type Rectangle struct{ W, H float64 }
type Circle struct{ R float64 }

// Rectangle と Circle がそれぞれ Name()/Area() を実装する。
// → どちらも自動的に Shape を満たす（"implements" と書く必要はない）。
func (r Rectangle) Name() string  { return "長方形" }
func (r Rectangle) Area() float64 { return r.W * r.H }
func (c Circle) Name() string     { return "円" }
func (c Circle) Area() float64    { return 3.14159 * c.R * c.R }

// 引数を Shape 型にすると、具体的な型を問わず同じように扱える。
func printArea(s Shape) {
	fmt.Printf("%s の面積 = %.2f\n", s.Name(), s.Area())
}

func main() {
	// 異なる型を同じ []Shape にまとめて扱える。
	shapes := []Shape{
		Rectangle{W: 3, H: 4},
		Circle{R: 5},
	}
	for _, s := range shapes {
		printArea(s)
	}
}
