/* eslint-disable */
// ブラウザ版（GitHub Pages）の Go 実行ワーカー（クラシック Worker）。
// Go の wasm ランタイム（wasm_exec.js）と yaegi をコンパイルした main.wasm を
// 読み込み、メインスレッドからのソースを yaegi で実行/整形して結果を返す。
//
// このワーカーで実行するため、無限ループが来てもメインスレッド（UI）は固まらず、
// タイムアウト時はメイン側が worker.terminate() でこのワーカーごと停止できる。
//
// パスは worker.js 自身の URL 基準で解決される（Pages の base 配下でも動く）。
importScripts('wasm_exec.js')

const go = new Go()

fetch('main.wasm')
  .then((r) => r.arrayBuffer())
  .then((bytes) => WebAssembly.instantiate(bytes, go.importObject))
  .then((result) => {
    go.run(result.instance) // main() が gopherjukuRun/Format を登録し select{} で常駐
    postMessage({ type: 'ready' })
  })
  .catch((err) => postMessage({ type: 'error', error: String(err) }))

self.onmessage = (e) => {
  const { id, kind, source } = e.data
  try {
    if (kind === 'run') {
      const r = self.gopherjukuRun(source)
      postMessage({ type: 'result', id, output: r.output, error: r.error })
    } else {
      const f = self.gopherjukuFormat(source)
      postMessage({ type: 'result', id, formatted: f.formatted, ok: f.ok })
    }
  } catch (err) {
    postMessage({ type: 'result', id, output: '', error: String(err) })
  }
}
