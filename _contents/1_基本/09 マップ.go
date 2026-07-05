// マップ
//
// マップはキーと値の対応表（連想配列）です。
package main

import "fmt"

func main() {
	// リテラルで初期化。
	age := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	// 追加・更新・参照。
	age["Carol"] = 28
	fmt.Println("Alice の年齢:", age["Alice"])

	// カンマ ok 記法：キーの有無を確認できる。
	if v, ok := age["Dave"]; ok {
		fmt.Println("Dave:", v)
	} else {
		fmt.Println("Dave は未登録")
	}

	// 削除。
	delete(age, "Bob")

	// fmt はマップをキー順にそろえて出力する（反復順は本来ランダム）。
	fmt.Println("現在のマップ:", age)
	fmt.Println("人数:", len(age))
}
