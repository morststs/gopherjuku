// ブラウザ版（GitHub Pages）の実行/整形マネージャ。
// public/worker.js（Go wasm ランナー）を Web Worker として起動し、
// 実行をタイムアウト付きで依頼する。タイムアウト時は worker を terminate して
// 再生成する（＝無限ループを確実に止める）。
//
// アセット（worker.js / wasm_exec.js / main.wasm）は Vite の BASE_URL 配下に置く。

const base = import.meta.env.BASE_URL

let worker = null
let readyPromise = null
let readyResolve = null
let seq = 0
const pending = new Map()

function spawn() {
  worker = new Worker(base + 'worker.js') // クラシック Worker（importScripts 用）
  worker.onmessage = (e) => {
    const d = e.data
    if (d.type === 'ready') {
      readyResolve && readyResolve()
    } else if (d.type === 'result') {
      const p = pending.get(d.id)
      if (p) {
        pending.delete(d.id)
        p(d)
      }
    } else if (d.type === 'error') {
      // wasm 読み込み失敗など。準備待ちを解除し、以降の呼び出しでエラーを返す。
      readyResolve && readyResolve()
    }
  }
}

// wasm のダウンロード・初期化を開始し、準備完了を待てる Promise を返す。
export function ensureReady() {
  if (!readyPromise) {
    readyPromise = new Promise((res) => {
      readyResolve = res
    })
    spawn()
  }
  return readyPromise
}

function call(kind, source, timeoutMs) {
  return new Promise((resolve) => {
    ensureReady().then(() => {
      const id = ++seq
      let settled = false
      const finish = (d) => {
        if (settled) return
        settled = true
        clearTimeout(timer)
        pending.delete(id)
        resolve(d)
      }
      const timer = setTimeout(() => {
        // 無限ループ等。ワーカーごと停止して作り直す。
        if (worker) worker.terminate()
        worker = null
        readyPromise = null
        readyResolve = null
        finish({ timedOut: true })
      }, timeoutMs)
      pending.set(id, finish)
      worker.postMessage({ id, kind, source })
    })
  })
}

export async function webRun(source) {
  const d = await call('run', source, 5000)
  if (d.timedOut) {
    return { success: false, output: '', error: '実行がタイムアウトしました（無限ループの可能性があります）。' }
  }
  return { success: !d.error, output: d.output || '', error: d.error || '' }
}

export async function webFormat(source) {
  const d = await call('format', source, 5000)
  if (d.timedOut) {
    return { success: false, formatted: source, error: '整形がタイムアウトしました。' }
  }
  return { success: !!d.ok, formatted: d.formatted != null ? d.formatted : source, error: '' }
}
