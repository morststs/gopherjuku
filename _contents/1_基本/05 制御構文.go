// 制御構文（if・for・switch）
//
// Go の分岐とループはとてもシンプルです。ループは for だけ。
package main

import "fmt"

func main() {
	// if：条件に丸括弧は不要。初期化文付きの if も書ける。
	if n := 7; n%2 == 0 {
		fmt.Println(n, "は偶数")
	} else {
		fmt.Println(n, "は奇数")
	}

	// for その1：カウンタ付きの基本形。
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("1..5 の合計:", sum)

	// for その2：条件だけ書けば while 相当。
	i := 1
	for i < 100 {
		i *= 2
	}
	fmt.Println("100 を超える最小の 2 の累乗:", i)

	// for その3：range でスライスを反復（添字と値を取り出す）。
	for idx, v := range []string{"a", "b", "c"} {
		fmt.Printf("[%d]=%s ", idx, v)
	}
	fmt.Println()

	// switch：break は不要。条件を書かない switch は if-else の連鎖に便利。
	for _, n := range []int{1, 2, 3} {
		switch {
		case n < 2:
			fmt.Println(n, "→ 小")
		case n < 3:
			fmt.Println(n, "→ 中")
		default:
			fmt.Println(n, "→ 大")
		}
	}
}
