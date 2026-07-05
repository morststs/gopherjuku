// 配列とスライス
//
// 配列は固定長、スライスは可変長。実務ではほぼスライスを使います。
package main

import "fmt"

func main() {
	// 配列：長さが型の一部（[3]int と [4]int は別の型）。
	var arr [3]int
	arr[0], arr[1], arr[2] = 10, 20, 30
	fmt.Println("配列:", arr, "長さ:", len(arr))

	// スライス：可変長。append で要素を追加できる。
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	fmt.Println("スライス:", s, "len:", len(s), "cap:", cap(s))

	// スライス式 s[low:high]（high は含まない）。
	fmt.Println("s[1:3]:", s[1:3])

	// make で長さ・容量を指定して作る。
	buf := make([]string, 0, 3)
	for _, w := range []string{"go", "is", "fun"} {
		buf = append(buf, w)
	}
	fmt.Println("buf:", buf)

	// スライスは参照的：元と一部を共有する点に注意。
	a := []int{1, 2, 3}
	b := a[:2]
	b[0] = 99
	fmt.Println("共有の例 a:", a, "b:", b) // a[0] も 99 になる
}
