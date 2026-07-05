// 型アサーションと型スイッチ
//
// インターフェース値の「中身の具体的な型」を取り出す方法です。
package main

import "fmt"

func main() {
	// interface{}（= any）はどんな値でも入る。
	var v any = "hello"

	// 型アサーション：v.(T)。カンマ ok 記法で安全に判定できる。
	if s, ok := v.(string); ok {
		fmt.Println("文字列だった。長さ:", len(s))
	}
	if _, ok := v.(int); !ok {
		fmt.Println("int ではない")
	}

	// 型スイッチ：複数の型をまとめて振り分ける。
	for _, x := range []any{42, "go", 3.14, true, []int{1, 2}} {
		fmt.Println(describe(x))
	}
}

func describe(i any) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return fmt.Sprintf("string: %q", v)
	case float64:
		return fmt.Sprintf("float64: %.2f", v)
	default:
		return fmt.Sprintf("その他: %T", v)
	}
}
