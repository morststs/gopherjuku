// ジェネリクス（型パラメータ）
//
// 型パラメータを使うと、型に依存しない汎用的な関数や型を書けます（Go 1.18 以降）。
package main

import "fmt"

// [T any] が型パラメータ。any 制約は「どんな型でも受け取れる」。
// 条件に合う要素だけを残す汎用フィルタ。T は呼び出しから推論される。
func Filter[T any](s []T, keep func(T) bool) []T {
	var r []T
	for _, v := range s {
		if keep(v) {
			r = append(r, v)
		}
	}
	return r
}

// 制約は自分で定義できる。~int は「基底型が int の型すべて」を表す。
type Number interface {
	~int | ~float64
}

// Number を満たす型なら何でも合計できる。
func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	// 同じ Filter が int にも string にも使える（型ごとに書かなくてよい）。
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("偶数:", evens)

	words := []string{"go", "gopher", "js", "rust"}
	long := Filter(words, func(w string) bool { return len(w) >= 3 })
	fmt.Println("3文字以上:", long)

	// 制約付きジェネリクス。int でも float64 でも同じ Sum で合計できる。
	fmt.Println("int の合計:", Sum(nums))
	fmt.Println("float の合計:", Sum([]float64{1.5, 2.5}))
}
