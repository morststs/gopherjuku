// 埋め込み（構造体・インターフェースの合成）
//
// Go は継承の代わりに「埋め込み」で機能を再利用します。
// 埋め込んだ型のフィールドやメソッドを、外側の型が自分のものとして使えます。
package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Describe() string {
	return a.Name + " です"
}

// Dog は Animal を「埋め込む」（フィールド名を書かない）。
type Dog struct {
	Animal // 埋め込み
	Breed  string
}

func (d Dog) Bark() string {
	return d.Name + ": ワン!" // 埋め込んだ Animal の Name に直接アクセスできる
}

func main() {
	d := Dog{
		Animal: Animal{Name: "レックス"},
		Breed:  "柴犬",
	}

	// Animal のメソッドを Dog がそのまま持っているように呼べる。
	fmt.Println(d.Describe())
	fmt.Println(d.Bark())
	fmt.Println("犬種:", d.Breed)

	// 埋め込んだフィールドは型名でも参照できる。
	fmt.Println("フルアクセス:", d.Animal.Name)
}
