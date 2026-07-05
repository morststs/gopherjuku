// strings — 文字列操作
//
// 文字列の検索・置換・分割・連結などをまとめて提供する定番パッケージ。
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, Gopher"

	fmt.Println("大文字:", strings.ToUpper(s))
	fmt.Println("含む(Go)?:", strings.Contains(s, "Go"))
	fmt.Println("置換:", strings.ReplaceAll(s, "o", "0"))
	fmt.Println("接頭辞(Hello)?:", strings.HasPrefix(s, "Hello"))

	// 分割と連結
	parts := strings.Split("a,b,c", ",")
	fmt.Printf("分割: %v（%d 個）\n", parts, len(parts))
	fmt.Println("連結:", strings.Join([]string{"x", "y", "z"}, "-"))

	// 前後の空白を除去
	fmt.Printf("トリム前後: %q\n", strings.TrimSpace("  padded  "))

	// Builder は連結を効率よく行う（+ の繰り返しより速い）。
	var b strings.Builder
	for i := 0; i < 3; i++ {
		b.WriteString("go")
	}
	fmt.Println("Builder:", b.String())
}
