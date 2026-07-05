// JSON — encoding/json による変換
//
// 構造体と JSON を相互に変換する。Web API や設定ファイルで頻出。
package main

import (
	"encoding/json"
	"fmt"
)

// 構造体タグ `json:"..."` で JSON 側のキー名を指定する。
type User struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Tags []string `json:"tags"`
}

func main() {
	u := User{Name: "Alice", Age: 30, Tags: []string{"go", "web"}}

	// 構造体 → JSON（Marshal）
	b, _ := json.Marshal(u)
	fmt.Println("JSON:", string(b))

	// インデント付きで整形出力
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(pretty))

	// JSON → 構造体（Unmarshal）。&v のようにポインタを渡す。
	var v User
	json.Unmarshal([]byte(`{"name":"Bob","age":25,"tags":["cli"]}`), &v)
	fmt.Printf("復元: %+v\n", v)
}
