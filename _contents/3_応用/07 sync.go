// sync — 並行処理の同期
//
// 複数の goroutine を安全に協調させる道具です。
// WaitGroup（完了待ち）, Mutex（排他制御）, Once（一度だけ実行）を扱います。
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Mutex：複数 goroutine から同じ変数を安全に更新する。
	var mu sync.Mutex
	counter := 0

	// Once：何度呼ばれても初期化を 1 回だけ実行する。
	var once sync.Once

	// WaitGroup：全 goroutine の完了を待つ。
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() { fmt.Println("初期化は一度だけ") })

			mu.Lock() // ロックしてから更新（競合を防ぐ）
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	// Mutex で守ったので、必ず 100 になる。
	fmt.Println("最終カウント:", counter)
}
