// time — 日付・時刻・期間
//
// 時刻の生成・加算・差分・整形を行います。
// （実行結果を一定にするため、ここでは固定の日時を使います。）
package main

import (
	"fmt"
	"time"
)

func main() {
	// 固定の日時を作る（time.Now() を使うと実行のたびに変わる）。
	t := time.Date(2026, time.July, 5, 14, 30, 0, 0, time.UTC)

	// Format は「参照時刻 2006-01-02 15:04:05」の並びで書式を指定する（Go 独特）。
	fmt.Println("整形:", t.Format("2006-01-02 15:04:05"))
	fmt.Println("和風:", t.Format("2006年01月02日 15時04分"))

	// Duration（期間）と加算。
	d := 90 * time.Minute
	fmt.Println("90分後:", t.Add(d).Format("15:04"))

	// 2 時刻の差分。
	t2 := time.Date(2026, time.July, 5, 12, 0, 0, 0, time.UTC)
	fmt.Println("差:", t.Sub(t2))

	// 各パーツの取り出し。
	fmt.Printf("年=%d 月=%d 日=%d 時=%d\n", t.Year(), t.Month(), t.Day(), t.Hour())
}
