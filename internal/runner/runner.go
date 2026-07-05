// Package runner は、ユーザーが入力した Go ソースを yaegi（純 Go 製の
// Go インタプリタ）で実行する処理を担う。
//
// yaegi はライブラリとして本体に static リンクされるため、外部の Go ツールチェーンや
// トランスパイラを一切必要としない。したがって単体の配布 exe だけで実行が完結する。
//
// 実行は「自分自身の実行ファイルを子プロセスとして起動し、標準入力でソースを渡す」
// 方式を採る。理由:
//   - 無限ループを含むコードでも、タイムアウトでプロセスごと確実に停止できる
//     （インプロセスの goroutine では強制停止できない）。
//   - GUI 本体（親プロセス）から実行を隔離できる。
//
// 子プロセス側の入口は RunChild（main.go から ChildFlag で分岐して呼ぶ）。
// テストや軽量用途向けに、同一プロセスで実行する Interpret も用意している。
package runner

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// ChildFlag は「インタプリタ子プロセスとして起動せよ」を示す第 1 引数。
const ChildFlag = "__interp"

// defaultTimeout は 1 回の実行を打ち切るまでの時間（無限ループ対策）。
const defaultTimeout = 5 * time.Second

// Result は実行結果。フロントエンド（Wails バインディング）へ JSON で渡される。
type Result struct {
	Success bool   `json:"success"` // 正常終了なら true
	Output  string `json:"output"`  // プログラムの標準出力（+標準エラー）
	Error   string `json:"error"`   // コンパイル/実行エラーやタイムアウトのメッセージ
}

// Runner は実行設定を保持する。
type Runner struct {
	timeout time.Duration
}

func New() *Runner {
	return &Runner{timeout: defaultTimeout}
}

// Run は自分自身を子プロセス（インタプリタモード）として起動し、source を
// 標準入力で渡して実行する。標準出力を集め、タイムアウト時はプロセスを kill する。
func (r *Runner) Run(source string) Result {
	exe, err := os.Executable()
	if err != nil {
		return Result{Error: "実行ファイルの特定に失敗しました: " + err.Error()}
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, exe, ChildFlag)
	cmd.Stdin = strings.NewReader(source)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	runErr := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return Result{
			Success: false,
			Output:  out.String(),
			Error:   fmt.Sprintf("実行がタイムアウトしました（%.0f 秒）。無限ループの可能性があります。", r.timeout.Seconds()),
		}
	}
	if runErr != nil {
		msg := strings.TrimSpace(errBuf.String())
		if msg == "" {
			msg = runErr.Error()
		}
		return Result{Success: false, Output: out.String(), Error: msg}
	}
	return Result{Success: true, Output: out.String()}
}

// Interpret は同一プロセスで yaegi 実行する（子プロセスを使わない）。
// テストや、タイムアウト隔離が不要な場面向け。
func Interpret(source string) Result {
	var out bytes.Buffer
	if err := interpret(source, &out); err != nil {
		return Result{Success: false, Output: out.String(), Error: strings.TrimSpace(err.Error())}
	}
	return Result{Success: true, Output: out.String()}
}

// RunChild は子プロセス側の入口。標準入力の Go ソースを yaegi で実行し、
// 出力を標準出力へ、エラーを標準エラーへ書いて終了コードで成否を伝える。
// main.go が os.Args[1] == ChildFlag のとき呼ぶ。
func RunChild() {
	defer func() {
		// yaegi 実行中の予期せぬ panic を拾ってエラーとして返す。
		if rec := recover(); rec != nil {
			fmt.Fprintf(os.Stderr, "実行時パニック: %v", rec)
			os.Exit(1)
		}
	}()

	src, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ソースの読み込みに失敗しました:", err)
		os.Exit(1)
	}
	if err := interpret(string(src), os.Stdout); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// interpret は yaegi インタプリタを生成し、プログラムの出力を out へ流して
// ソースを評価する。package main / func main を含む完全なソースを渡すと
// yaegi が main を実行する。
func interpret(source string, out io.Writer) error {
	i := interp.New(interp.Options{Stdout: out, Stderr: out})
	// 標準ライブラリのシンボルを公開する（fmt, strings, math, ... を import 可能に）。
	if err := i.Use(stdlib.Symbols); err != nil {
		return err
	}
	_, err := i.Eval(source)
	return err
}
