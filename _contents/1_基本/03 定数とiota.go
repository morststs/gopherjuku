// 定数と iota
//
// const で定数を宣言します。iota を使うと連番の定数を簡潔に定義できます。
package main

import "fmt"

// 単純な定数
const Pi = 3.14159
const Greeting = "こんにちは"

// iota は const ブロック内で 0 から始まり、行ごとに 1 増える。
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
)

// iota を式に使うと、ビットフラグや単位も簡潔に書ける。
const (
	_  = iota             // 0 は捨てる
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
	GB                    // 1 << 30
)

func main() {
	fmt.Println(Greeting, "Pi =", Pi)
	fmt.Println("曜日:", Sunday, Monday, Tuesday, Wednesday)
	fmt.Println("KB, MB, GB:", KB, MB, GB)
}
