// 関数
//
// Go の関数は「複数の戻り値」「可変長引数」「名前付き戻り値」を持てます。
package main

import "fmt"

// 複数の戻り値。エラーを 2 つ目に返すのが Go の定石。
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("0 では割れません")
	}
	return a / b, nil
}

// 可変長引数（...int）。呼び出し側は任意個の引数を渡せる。
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 名前付き戻り値。return だけで名前付きの値が返る。
func split(total int) (a, b int) {
	a = total * 4 / 9
	b = total - a
	return
}

func main() {
	q, err := divide(10, 3)
	fmt.Println("10/3 =", q, "err:", err)

	_, err = divide(1, 0)
	fmt.Println("1/0 err:", err)

	fmt.Println("合計:", sum(1, 2, 3, 4, 5))

	x, y := split(18)
	fmt.Println("split(18) =", x, y)
}
