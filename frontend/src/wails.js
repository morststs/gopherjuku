// バックエンド呼び出しのラッパー。2 つの動作モードに対応する:
//
//   1. デスクトップ（Wails）: window.go.main.App.* を呼ぶ。
//   2. ブラウザ（GitHub Pages・静的配信）: バックエンドが無いので、
//      レッスンは静的 JSON（lessons.json）から、実行/整形は Go を wasm に
//      した Web Worker（webRunner.js）から供給する。
//
// どちらのモードかは window.go の有無で判定する。

const app = () => (typeof window !== 'undefined' ? window.go?.main?.App : undefined)

// ---- ブラウザ版: レッスンの静的データ ----
let lessonsPromise = null
function loadLessons() {
  if (!lessonsPromise) {
    lessonsPromise = fetch(import.meta.env.BASE_URL + 'lessons.json').then((r) => r.json())
  }
  return lessonsPromise
}

/** 左ペイン用のカテゴリー/題ツリーを取得する。 */
export async function fetchTree() {
  const b = app()
  if (b?.Tree) return b.Tree()
  const data = await loadLessons()
  return (data.categories || []).map((c) => ({
    name: c.name,
    lessons: (c.lessons || []).map((l) => ({ category: l.category, title: l.title, path: l.path })),
  }))
}

/** レッスンの Go ソースを取得する。 */
export async function fetchSource(path) {
  const b = app()
  if (b?.Source) return b.Source(path)
  const data = await loadLessons()
  for (const c of data.categories || []) {
    for (const l of c.lessons || []) {
      if (l.path === path) return l.source
    }
  }
  return ''
}

/**
 * Go ソースを実行し、標準出力とエラーを返す。
 * @returns {Promise<{success:boolean, output:string, error:string}>}
 */
export async function run(source) {
  const b = app()
  if (b?.Run) return b.Run(source)
  const { webRun } = await import('./webRunner.js')
  return webRun(source)
}

/**
 * Go ソースを gofmt 整形する（Shift+Alt+F から呼ばれる）。
 * @returns {Promise<{success:boolean, formatted:string, error:string}>}
 */
export async function format(source) {
  const b = app()
  if (b?.Format) return b.Format(source)
  const { webFormat } = await import('./webRunner.js')
  return webFormat(source)
}

/**
 * ブラウザ版で、実行環境（wasm）とレッスンの読み込みを先行開始する。
 * デスクトップ版では何もしない。
 */
export function prewarm() {
  if (app()) return
  loadLessons()
  import('./webRunner.js').then((m) => m.ensureReady())
}
