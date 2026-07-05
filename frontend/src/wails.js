// バックエンド（Go / Wails）呼び出しのラッパー。
//
// Wails は実行時に window.go.main.App.* を注入する。ブラウザ単体（wails なし）の
// UI 開発時は window.go が無いので、簡易フォールバックで最低限動くようにしている。
// これにより wailsjs の自動生成バインディングに依存せず vite build が通る。

const app = () => (typeof window !== 'undefined' ? window.go?.main?.App : undefined)

/** 左ペイン用のカテゴリー/題ツリーを取得する。 */
export async function fetchTree() {
  const b = app()
  if (b?.Tree) return b.Tree()
  // フォールバック（ブラウザ単体）
  return [
    { name: '基本', lessons: [{ category: '基本', title: 'ハローワールド', path: '基本/ハローワールド' }] },
  ]
}

/** レッスンの Go ソースを取得する。 */
export async function fetchSource(path) {
  const b = app()
  if (b?.Source) return b.Source(path)
  return [
    'package main',
    '',
    'import "fmt"',
    '',
    'func main() {',
    '\tfmt.Println("Hello, Gopherjuku!")',
    '}',
    '',
  ].join('\n')
}

/**
 * Go ソースを yaegi インタプリタ（バックエンドの子プロセス）で実行し、
 * 標準出力とエラーを返す。
 * @returns {Promise<{success:boolean, output:string, error:string}>}
 */
export async function run(source) {
  const b = app()
  if (b?.Run) return b.Run(source)
  return {
    success: false,
    output: '',
    error: '（ブラウザ単体モード）バックエンド未接続のため実行できません。`wails dev` またはビルドした exe で起動してください。',
  }
}

/**
 * Go ソースを gofmt 整形する（エディタの Shift+Alt+F から呼ばれる）。
 * @returns {Promise<{success:boolean, formatted:string, error:string}>}
 */
export async function format(source) {
  const b = app()
  if (b?.Format) return b.Format(source)
  // ブラウザ単体モードでは整形せず、そのまま返す。
  return { success: false, formatted: source, error: '（ブラウザ単体モード）整形はバックエンドが必要です。' }
}
