<script>
  // ルート: 3 ペイン（左=レッスン, 中央=エディタ, 右=実行結果）+ 実行フロー。
  import { onMount } from 'svelte'
  import LessonTree from './LessonTree.svelte'
  import Editor from './Editor.svelte'
  import Output from './Output.svelte'
  import { fetchTree, fetchSource, run as runBackend, prewarm } from './wails.js'

  let categories = $state([])
  let currentPath = $state(null)
  let source = $state('')
  let output = $state('')
  let status = $state('idle') // idle | running | error
  let message = $state('')

  // 左右ペインの幅（中央のスプリッターでドラッグ調整）。
  let leftW = $state(240)
  let rightW = $state(380)

  // 狭い画面（スマホ縦など）では 3 ペイン横並びが破綻するため、タブ切り替えに変える。
  let narrow = $state(false)
  let tab = $state('editor') // narrow 時に表示するペイン: lessons | editor | output

  const busy = $derived(status === 'running')

  // 画面幅を監視して narrow を切り替える。
  $effect(() => {
    const mq = window.matchMedia('(max-width: 720px)')
    const update = () => (narrow = mq.matches)
    update()
    mq.addEventListener('change', update)
    return () => mq.removeEventListener('change', update)
  })

  onMount(async () => {
    prewarm() // ブラウザ版: 実行環境(wasm)とレッスンの読み込みを先行開始
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
    if (narrow) tab = 'editor' // レッスンを選んだらエディタへ移動
  }

  async function run() {
    if (narrow) tab = 'output' // 実行したら結果タブへ移動
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

<div class="app" class:narrow>
  <!-- 広い画面: 3 ペイン横並び（実行ボタンは実行結果ペインのヘッダ）。
       狭い画面: タブで 1 ペインずつ全画面表示（実行ボタンはタブバー）。 -->
  {#if narrow}
    <nav class="tabbar">
      <button class="tab" class:active={tab === 'lessons'} onclick={() => (tab = 'lessons')}>レッスン</button>
      <button class="tab" class:active={tab === 'editor'} onclick={() => (tab = 'editor')}>エディタ</button>
      <button class="tab" class:active={tab === 'output'} onclick={() => (tab = 'output')}>実行結果</button>
      <span class="tab-spacer"></span>
      <button class="tab-run" onclick={run} disabled={busy}>▶ 実行</button>
    </nav>
  {/if}

  <div class="panes" style={narrow ? '' : `grid-template-columns: ${leftW}px 6px 1fr 6px ${rightW}px;`}>
    <aside class="pane left" class:hidden={narrow && tab !== 'lessons'}>
      <LessonTree {categories} {currentPath} onselect={select} />
    </aside>

    {#if !narrow}
      <div class="gutter" onpointerdown={(e) => startDrag('left', e)}></div>
    {/if}

    <main class="pane center" class:hidden={narrow && tab !== 'editor'}>
      <Editor bind:value={source} />
    </main>

    {#if !narrow}
      <div class="gutter" onpointerdown={(e) => startDrag('right', e)}></div>
    {/if}

    <section class="pane right" class:hidden={narrow && tab !== 'output'}>
      <Output {output} {status} {message} onrun={run} {busy} showRun={!narrow} />
    </section>
  </div>
</div>

<style>
  .app {
    height: 100%;
    display: flex;
    flex-direction: column;
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

  /* --- 狭い画面（スマホ縦など）: タブ切り替えで 1 ペインずつ全画面 --- */
  .tabbar {
    display: flex;
    align-items: center;
    height: 48px;
    background: #323233;
    border-bottom: 1px solid #444;
    flex: none;
  }
  .tab {
    height: 100%;
    padding: 0 14px;
    background: none;
    border: none;
    border-bottom: 2px solid transparent;
    color: #cbd5e1;
    font-size: 14px;
    cursor: pointer;
  }
  .tab.active {
    color: #fff;
    border-bottom-color: #16a34a;
  }
  .tab-spacer {
    flex: 1;
  }
  .tab-run {
    margin-right: 8px;
    background: #16a34a;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding: 8px 16px;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
  }
  .tab-run:disabled {
    opacity: 0.5;
  }
  .app.narrow .panes {
    display: block;
    position: relative;
  }
  .app.narrow .pane {
    height: 100%;
  }
  .pane.hidden {
    display: none;
  }
</style>
