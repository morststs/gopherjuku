package gofmt

import "testing"

func TestSourceFormats(t *testing.T) {
	// インデントや空白が崩れたソースを渡す。
	in := "package main\nimport \"fmt\"\nfunc main(){fmt.Println( \"hi\" )}\n"
	res := Source(in)
	if !res.Success {
		t.Fatalf("整形成功を期待したが失敗: %s", res.Error)
	}
	// gofmt はタブインデント・括弧内の余分な空白除去などを行う。
	want := "package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"hi\") }\n"
	if res.Formatted != want {
		t.Fatalf("整形結果が不一致:\n got: %q\nwant: %q", res.Formatted, want)
	}
}

func TestSourceSyntaxError(t *testing.T) {
	in := "package main\nfunc main() {\n" // 閉じ括弧なし
	res := Source(in)
	if res.Success {
		t.Fatal("失敗を期待したが成功した")
	}
	if res.Formatted != in {
		t.Fatal("失敗時は入力をそのまま返すべき")
	}
	if res.Error == "" {
		t.Fatal("エラーメッセージを期待したが空")
	}
}
