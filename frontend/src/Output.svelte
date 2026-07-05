<script>
  // 右ペイン: 実行結果（標準出力・エラー・ステータス）。ヘッダに実行ボタンを持つ。
  // showRun=false のとき（狭い画面ではタブバー側に実行ボタンを置く）ヘッダの実行ボタンは隠す。
  let { output = '', status = 'idle', message = '', onrun, busy = false, showRun = true } = $props()

  // 出力テキストの文字サイズ（Ctrl+マウスホイールで変更）。
  let fontSize = $state(13)

  function onWheel(e) {
    if (!e.ctrlKey) return
    e.preventDefault() // ブラウザ/WebView 全体のズームを抑止し、出力だけ拡大縮小する
    const next = fontSize + (e.deltaY < 0 ? 1 : -1)
    fontSize = Math.max(8, Math.min(40, next))
  }
</script>

<div class="output" onwheel={onWheel}>
  <div class="out-head">
    <span class="title">実行結果</span>
    <span class="spacer"></span>
    {#if showRun}
      <button class="run" onclick={onrun} disabled={busy}>▶ 実行</button>
    {/if}
  </div>

  {#if message}
    <div
      class="banner"
      class:error={status === 'error'}
      class:busy={status === 'running'}
      style="font-size: {fontSize}px"
    >
      {message}
    </div>
  {/if}

  <pre class="console" style="font-size: {fontSize}px">{output}</pre>
</div>

<style>
  .output {
    height: 100%;
    display: flex;
    flex-direction: column;
    background: #1e1e1e;
    border-left: 1px solid #333;
  }
  .out-head {
    display: flex;
    align-items: center;
    gap: 12px;
    height: 44px;
    padding: 0 12px;
    background: #323233;
    border-bottom: 1px solid #444;
    flex: none;
  }
  .out-head .title {
    font-size: 11px;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: #9ca3af;
  }
  .out-head .spacer {
    flex: 1;
  }
  .out-head .run {
    background: #16a34a;
    color: #fff;
    border: none;
    border-radius: 6px;
    padding: 6px 16px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
  }
  .out-head .run:hover:not(:disabled) {
    background: #15803d;
  }
  .out-head .run:disabled {
    opacity: 0.5;
    cursor: default;
  }
  .banner {
    padding: 8px 12px;
    background: #2d2d30;
    color: #cbd5e1;
    white-space: pre-wrap;
  }
  .banner.busy {
    color: #93c5fd;
  }
  .banner.error {
    background: #3b1d1d;
    color: #fca5a5;
  }
  .console {
    margin: 0;
    padding: 12px;
    flex: 1;
    overflow: auto;
    font-family: 'Cascadia Code', 'Consolas', 'Menlo', monospace;
    line-height: 1.5;
    white-space: pre-wrap;
    word-break: break-word;
    color: #d4d4d4;
  }
</style>
