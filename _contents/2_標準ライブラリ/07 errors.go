// errors — エラーの生成と判定
//
// エラー値の作り方、独自エラー型、ラップと一致判定を扱います。
package main

import (
	"errors"
	"fmt"
)

// 独自のエラー型：Error() string を実装すれば error として使える。
type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s が不正です", e.Field)
}

// 比較の目印になる sentinel エラー。
var ErrEmpty = errors.New("空です")

func validate(name string) error {
	if name == "" {
		// %w でラップすると、元のエラーを保持したまま文脈を付けられる。
		return fmt.Errorf("検証失敗: %w", ErrEmpty)
	}
	return nil
}

func main() {
	// errors.New：シンプルなエラー。
	fmt.Println("New:", errors.New("何かに失敗"))

	// 独自エラー型。Error() の文言が使われる。
	var verr error = &ValidationError{Field: "name"}
	fmt.Println("独自型:", verr)

	// ラップしたエラーは errors.Is で元の sentinel と一致判定できる。
	err := validate("")
	fmt.Println("検証:", err)
	if errors.Is(err, ErrEmpty) {
		fmt.Println("  → 原因は ErrEmpty")
	}
}
