<script>
  // ルート: 3 ペイン（左=レッスン, 中央=エディタ, 右=実行結果）+ 実行フロー。
  import { onMount } from 'svelte'
  import LessonTree from './LessonTree.svelte'
  import Editor from './Editor.svelte'
  import Output from './Output.svelte'
  import { fetchTree, fetchSource, run as runBackend } from './wails.js'

  let categories = $state([])
  let currentPath = $state(null)
  let source = $state('')
  let output = $state('')
  let status = $state('idle') // idle | running | error
  let message = $state('')

  // 左右ペインの幅（中央のスプリッターでドラッグ調整）。
  let leftW = $state(240)
  let rightW = $state(380)

  const busy = $derived(status === 'running')

  onMount(async () => {
    categories = await fetchTree()
    const first = categories[0]?.lessons?.[0]
    if (first) await select(first.path)
  })

  async function select(path) {
    currentPath = path
    source = await fetchSource(path)
    output = ''
    message = ''
    status = 'idle'
  }

  async function run() {
    status = 'running'
    message = '実行中…'
    output = ''

    const res = await runBackend(source)
    output = res.output || ''
    if (!res.success) {
      status = 'error'
      message = res.error || '実行に失敗しました。'
    } else {
      status = 'idle'
      message = ''
    }
  }

  // スプリッターのドラッグ処理。
  function startDrag(which, e) {
    e.preventDefault()
    const startX = e.clientX
    const startLeft = leftW
    const startRight = rightW
    const move = (ev) => {
      const dx = ev.clientX - startX
      if (which === 'left') {
        leftW = Math.max(160, Math.min(480, startLeft + dx))
      } else {
        rightW = Math.max(200, Math.min(680, startRight - dx))
      }
    }
    const up = () => {
      window.removeEventListener('pointermove', move)
      window.removeEventListener('pointerup', up)
    }
    window.addEventListener('pointermove', move)
    window.addEventListener('pointerup', up)
  }
</script>

<div class="app">
  <header class="topbar">
    <!-- ツールバーは実行ボタンのみ（アプリ名は OS のタイトルバーに表示）。 -->
    <span class="spacer"></span>
    <button class="run" onclick={run} disabled={busy}>▶ 実行</button>
  </header>

  <div class="panes" style="grid-template-columns: {leftW}px 6px 1fr 6px {rightW}px;">
    <aside class="pane left">
      <LessonTree {categories} {currentPath} onselect={select} />
    </aside>

    <div class="gutter" onpointerdown={(e) => startDrag('left', e)}></div>

    <main class="pane center">
      <Editor bind:value={source} />
    </main>

    <div class="gutter" onpointerdown={(e) => startDrag('right', e)}></div>

    <section class="pane right">
      <Output {output} {status} {message} />
    </section>
  </div>
</div>

<style>
  .app {
    height: 100%;
    display: flex;
    flex-direction: column;
  }
  .topbar {
    display: flex;
    align-items: center;
    gap: 12px;
    height: 44px;
    padding: 0 14px;
    background: #323233;
    border-bottom: 1px solid #444;
    flex: none;
  }
  .spacer {
    flex: 1;
  }
  .run {
    background: #16a34a;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding: 6px 16px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
  }
  .run:hover:not(:disabled) {
    background: #15803d;
  }
  .run:disabled {
    opacity: 0.5;
    cursor: default;
  }
  .panes {
    flex: 1;
    display: grid;
    min-height: 0;
  }
  .pane {
    min-width: 0;
    overflow: hidden;
  }
  .gutter {
    background: #333;
    cursor: col-resize;
  }
  .gutter:hover {
    background: #0e639c;
  }
</style>
