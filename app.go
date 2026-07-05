package main

import (
	"context"

	"gopherjuku/internal/gofmt"
	"gopherjuku/internal/lessons"
	"gopherjuku/internal/runner"
)

// App は Wails にバインドされるメインの構造体。
// フロントエンド（Svelte）からは window.go.main.App.XXX として呼び出せる。
type App struct {
	ctx     context.Context
	runner  *runner.Runner
	lessons *lessons.Service
}

// NewApp は依存（yaegi ランナー・教材リポジトリ）を組み立てて App を返す。
func NewApp(ls *lessons.Service) *App {
	return &App{
		runner:  runner.New(),
		lessons: ls,
	}
}

// startup は Wails 起動時に呼ばれ、コンテキストを保持する。
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Tree は左ペイン用のカテゴリー/題ツリーを返す。
func (a *App) Tree() []lessons.Category {
	return a.lessons.Tree()
}

// Source は指定レッスンの Go ソースを返す。
func (a *App) Source(path string) (string, error) {
	return a.lessons.Source(path)
}

// Run は中央エディターの Go ソースを yaegi インタプリタ（子プロセス）で実行し、
// 標準出力とエラーを返す。結果はそのまま右ペインに表示される。
func (a *App) Run(source string) runner.Result {
	return a.runner.Run(source)
}

// Format は Go ソースを gofmt 整形して返す。エディタの「Format Document」
// (Shift+Alt+F) から呼ばれる。
func (a *App) Format(source string) gofmt.Result {
	return gofmt.Source(source)
}
