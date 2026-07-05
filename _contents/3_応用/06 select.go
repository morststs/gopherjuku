// select — 複数チャネルの待ち受け
//
// select は複数の channel 操作のうち、準備ができたものを 1 つ選んで実行します。
// タイムアウトやキャンセルの実装によく使います。
package main

import (
	"fmt"
	"time"
)

func main() {
	fast := make(chan string)
	slow := make(chan string)

	// 2 つの goroutine が別々のタイミングで結果を送る。
	go func() { time.Sleep(20 * time.Millisecond); fast <- "fast の結果" }()
	go func() { time.Sleep(200 * time.Millisecond); slow <- "slow の結果" }()

	// 先に届いた方を受け取る。両方来るまで 2 回ループ。
	for i := 0; i < 2; i++ {
		select {
		case v := <-fast:
			fmt.Println("受信:", v)
		case v := <-slow:
			fmt.Println("受信:", v)
		case <-time.After(time.Second):
			fmt.Println("タイムアウト")
		}
	}

	// time.After を使った単独のタイムアウト例。
	done := make(chan bool)
	go func() { time.Sleep(10 * time.Millisecond); done <- true }()
	select {
	case <-done:
		fmt.Println("処理完了")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("時間切れ")
	}
}
