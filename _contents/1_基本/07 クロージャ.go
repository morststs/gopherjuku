// クロージャ
//
// 関数は値として扱えます。関数が外側の変数を「捕まえて」保持したものがクロージャです。
package main

import "fmt"

// 呼ぶたびに 1 増える「カウンタ」を返す。count は各クロージャに固有。
func counter() func() int {
	count := 0
	return func() int {
		count++ // 外側の count を捕まえて更新し続ける
		return count
	}
}

func main() {
	// 関数を変数に入れて使える。
	double := func(n int) int { return n * 2 }
	fmt.Println("double(21) =", double(21))

	// counter() が返すクロージャは、それぞれ独立した count を持つ。
	c1 := counter()
	c2 := counter()
	fmt.Println("c1:", c1(), c1(), c1()) // 1 2 3
	fmt.Println("c2:", c2())             // 1（c1 とは別）
}
