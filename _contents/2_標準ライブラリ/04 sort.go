// sort — 並べ替え
//
// スライスを昇順に並べたり、独自の基準で並べ替えたりできる。
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 基本型は専用関数が用意されている。
	ints := []int{3, 1, 4, 1, 5, 9, 2}
	sort.Ints(ints)
	fmt.Println("昇順:", ints)

	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println("辞書順:", strs)

	// sort.Slice なら任意の基準で並べ替えられる（ここでは文字数の降順）。
	words := []string{"go", "gopher", "js", "rust"}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	fmt.Println("長さの降順:", words)

	// 昇順に並んだスライスなら二分探索できる。
	pos := sort.SearchInts(ints, 4)
	fmt.Printf("値 4 は位置 %d\n", pos)
}
