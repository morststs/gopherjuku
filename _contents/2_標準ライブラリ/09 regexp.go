// regexp — 正規表現
//
// パターンによる検索・抽出・置換を行います。
// パターンはバッククォート文字列で書くと、バックスラッシュをそのまま書けて便利です。
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// MustCompile はパターンをコンパイルする（不正なら panic）。
	re := regexp.MustCompile(`\d+`) // 1 個以上の数字

	fmt.Println("数字を含む?:", re.MatchString("abc123"))
	fmt.Println("最初の一致:", re.FindString("a12b345"))
	fmt.Println("全部の一致:", re.FindAllString("a12b345c6", -1))
	fmt.Println("置換:", re.ReplaceAllString("a1b2c3", "#"))

	// グループ () で部分を取り出す。
	mail := regexp.MustCompile(`(\w+)@(\w+)`)
	m := mail.FindStringSubmatch("user@example")
	// m[0]=全体, m[1]=1つ目のグループ, m[2]=2つ目
	fmt.Printf("全体=%s ユーザー=%s ドメイン=%s\n", m[0], m[1], m[2])
}
