// math と乱数（math / math/rand）
//
// 数学関数と擬似乱数を使います。
// （結果を一定にするため、乱数は固定シードで生成します。）
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// math：定数と関数。
	fmt.Printf("Pi=%.5f  E=%.5f\n", math.Pi, math.E)
	fmt.Printf("Sqrt(2)=%.4f  Pow(2,10)=%.0f\n", math.Sqrt(2), math.Pow(2, 10))
	fmt.Println("Abs(-5), Ceil(2.1), Floor(2.9):", math.Abs(-5), math.Ceil(2.1), math.Floor(2.9))
	fmt.Println("Max, Min:", math.Max(3, 7), math.Min(3, 7))

	// math/rand：固定シードにすると毎回同じ列になる（再現性のため）。
	r := rand.New(rand.NewSource(42))
	fmt.Println("0..99 の乱数:", r.Intn(100), r.Intn(100), r.Intn(100))

	// シャッフルも rand で行える。
	nums := []int{1, 2, 3, 4, 5}
	r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	fmt.Println("シャッフル:", nums)
}
