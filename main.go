package main

import (
	"embed"
	"os"

	"gopherjuku/internal/lessons"
	"gopherjuku/internal/runner"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// フロントエンドのビルド成果物を同梱する（wails build が frontend/dist を生成する）。
//
//go:embed all:frontend/dist
var assets embed.FS

// 教材（コードの題）を同梱する。追加は _contents/<カテゴリー>/<題>.go に置くだけ。
// ディレクトリ名を "_" で始めているのは、埋め込む .go ファイルを Go のビルド対象
// （パッケージ）として走査させないため。go ツールは "_" 始まりのディレクトリを無視し、
// 一方で go:embed の all: 接頭辞は "_" 始まりでも取り込む。
//
//go:embed all:_contents
var contentsFS embed.FS

func main() {
	// インタプリタ子プロセスモード: 標準入力の Go ソースを yaegi で実行して終了する。
	// GUI 本体（親）が自分自身をこのモードで起動し、タイムアウト時に kill できる。
	// GUI（wails.Run）に入る前に必ずここで分岐すること。
	if len(os.Args) > 1 && os.Args[1] == runner.ChildFlag {
		runner.RunChild()
		return
	}

	app := NewApp(lessons.New(contentsFS))

	err := wails.Run(&options.App{
		Title:  "Gopherjuku",
		Width:  1200,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// アプリの背景（body #1e1e1e）に合わせる。
		BackgroundColour: &options.RGBA{R: 30, G: 30, B: 30, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		// Windows のタイトルバーをアプリの背景（body #1e1e1e = RGB(30,30,30)）と
		// 同じ色にする（sirusita と同じ CustomTheme 方式）。
		Windows: &windows.Options{
			Theme: windows.Dark,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:          windows.RGB(30, 30, 30),
				DarkModeTitleBarInactive:  windows.RGB(30, 30, 30),
				DarkModeTitleText:         windows.RGB(229, 231, 235),
				DarkModeTitleTextInactive: windows.RGB(150, 150, 150),
				DarkModeBorder:            windows.RGB(30, 30, 30),
				DarkModeBorderInactive:    windows.RGB(30, 30, 30),
			},
		},
		Linux: &linux.Options{
			ProgramName: "Gopherjuku",
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
