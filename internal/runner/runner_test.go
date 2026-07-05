package runner

import (
	"os"
	"strings"
	"testing"
	"time"
)

// TestMain は、テストバイナリが子プロセス（ChildFlag 付き）として起動された場合に
// インタプリタとして振る舞えるようにする。これにより Run（親→子プロセス→タイムアウト）
// の経路を、実アプリのビルドなしに検証できる。
func TestMain(m *testing.M) {
	if len(os.Args) > 1 && os.Args[1] == ChildFlag {
		RunChild()
		return
	}
	os.Exit(m.Run())
}

func TestInterpretHelloWorld(t *testing.T) {
	res := Interpret("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"hi\")\n}\n")
	if !res.Success {
		t.Fatalf("成功を期待したが失敗: %s", res.Error)
	}
	if res.Output != "hi\n" {
		t.Fatalf("出力が不一致: %q", res.Output)
	}
}

func TestInterpretError(t *testing.T) {
	res := Interpret("package main\n\nfunc main() {\n\tundefinedFn()\n}\n")
	if res.Success {
		t.Fatal("失敗を期待したが成功した")
	}
	if strings.TrimSpace(res.Error) == "" {
		t.Fatal("エラーメッセージを期待したが空")
	}
}

func TestRunViaChildProcess(t *testing.T) {
	r := New()
	res := r.Run("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfor i := 1; i <= 3; i++ {\n\t\tfmt.Println(i)\n\t}\n}\n")
	if !res.Success {
		t.Fatalf("成功を期待したが失敗: %s", res.Error)
	}
	if res.Output != "1\n2\n3\n" {
		t.Fatalf("出力が不一致: %q", res.Output)
	}
}

func TestRunTimeout(t *testing.T) {
	r := New()
	r.timeout = 800 * time.Millisecond // 無限ループを短時間で打ち切る
	res := r.Run("package main\n\nfunc main() {\n\tfor {\n\t}\n}\n")
	if res.Success {
		t.Fatal("タイムアウト（失敗）を期待したが成功した")
	}
	if !strings.Contains(res.Error, "タイムアウト") {
		t.Fatalf("タイムアウトのメッセージを期待したが: %s", res.Error)
	}
}
