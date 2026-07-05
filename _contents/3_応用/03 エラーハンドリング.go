// エラーハンドリング
//
// Go はエラーを「戻り値」で表現します（例外機構ではない）。
// 関数は error を返し、呼び出し側が明示的に確認するのが基本です。
package main

import (
	"errors"
	"fmt"
)

// sentinel エラー：比較の目印としてあらかじめ定義したエラー値。
var ErrNotFound = errors.New("見つかりません")

func find(id int) error {
	if id != 1 {
		// %w でラップすると、元のエラーを保持したまま文脈を追加できる。
		return fmt.Errorf("id=%d の検索に失敗: %w", id, ErrNotFound)
	}
	return nil
}

func main() {
	for _, id := range []int{1, 2} {
		err := find(id)
		if err == nil {
			fmt.Printf("id=%d: OK\n", id)
			continue
		}
		fmt.Printf("id=%d: %v\n", id, err)

		// errors.Is で、ラップされた元エラーと一致するか判定できる。
		if errors.Is(err, ErrNotFound) {
			fmt.Println("  → 原因は ErrNotFound")
		}
	}
}
