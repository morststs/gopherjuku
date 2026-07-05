// defer / panic / recover
//
// defer は「関数を抜けるときに実行」する予約、panic は異常終了、
// recover は panic を捕まえて復帰する仕組みです。
package main

import "fmt"

func main() {
	// defer は後入れ先出し（LIFO）で、関数の最後にまとめて実行される。
	// 後片付け（ファイルを閉じる等）に使う。
	fmt.Println("--- defer の順序 ---")
	demoDefer()

	// panic を recover で捕まえ、エラーとして返す例。
	fmt.Println("--- recover ---")
	err := safeDivide(10, 0)
	fmt.Println("結果:", err)
	fmt.Println("プログラムは継続している")
}

func demoDefer() {
	defer fmt.Println("defer 1（最後に出る）")
	defer fmt.Println("defer 2")
	fmt.Println("本体の処理")
}

// panic しても recover で復帰し、error に変換して返す。
func safeDivide(a, b int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("復帰しました: %v", r)
		}
	}()
	if b == 0 {
		panic("ゼロ除算")
	}
	fmt.Println(a / b)
	return nil
}
