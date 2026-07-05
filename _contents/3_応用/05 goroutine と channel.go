// goroutine と channel
//
// goroutine は軽量な並行実行の単位、channel はその間でデータを受け渡す通信路です。
// go キーワードで関数を並行に走らせます。
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := []int{2, 3, 4}

	// バッファ付き channel に各 goroutine の結果を集める。
	results := make(chan int, len(jobs))

	// WaitGroup で「全 goroutine の完了」を待つ。
	var wg sync.WaitGroup
	for _, n := range jobs {
		wg.Add(1)
		go func(x int) { // x を引数で渡すのが定石（ループ変数の共有を避ける）
			defer wg.Done()
			results <- x * x
		}(n)
	}

	wg.Wait()      // 全部終わるまで待つ
	close(results) // もう送らないので閉じる

	// 閉じた channel は range で最後まで読み切れる。
	sum := 0
	for r := range results {
		sum += r
	}
	// 実行順序は毎回変わるが、合計は一定：4 + 9 + 16 = 29
	fmt.Println("二乗の合計:", sum)
}
