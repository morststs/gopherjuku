// strconv — 文字列と数値の変換
//
// 文字列 ↔ 数値・真偽値の変換を行います。
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 文字列 → 数値。失敗するとエラーを返す。
	n, err := strconv.Atoi("42")
	fmt.Println("Atoi(\"42\"):", n, err)

	_, err = strconv.Atoi("abc")
	fmt.Println("Atoi(\"abc\") のエラー:", err)

	// 数値 → 文字列。
	s := strconv.Itoa(2026)
	fmt.Println("Itoa(2026):", s, "型は string")

	// 浮動小数点・真偽値。
	f, _ := strconv.ParseFloat("3.14", 64)
	b, _ := strconv.ParseBool("true")
	fmt.Println("ParseFloat, ParseBool:", f, b)

	// 基数を指定した変換（2 進数の文字列 → 10 進数）。
	v, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Println("2進 1010 =", v)

	// Quote：エスケープ済みの引用符付き文字列。
	fmt.Println("Quote:", strconv.Quote("行1\n行2"))
}
