<script>
  // 中央ペイン: Monaco エディタ（Go）。value を親と双方向バインドする。
  import { onMount, onDestroy } from 'svelte'
  import { monaco } from './monaco.js'

  let { value = $bindable('') } = $props()

  let container
  let editor
  let applyingProp = false // prop→エディタ反映中の変更イベントを無視するフラグ

  onMount(() => {
    editor = monaco.editor.create(container, {
      value,
      language: 'go',
      theme: 'vs-dark',
      automaticLayout: true,
      minimap: { enabled: false },
      fontSize: 14,
      tabSize: 4,
      scrollBeyondLastLine: false,
      // Ctrl + マウスホイールで文字サイズを拡大縮小する。
      mouseWheelZoom: true,
    })
    editor.onDidChangeModelContent(() => {
      if (applyingProp) return
      value = editor.getValue()
    })
  })

  // 親が value を差し替えたとき（レッスン切替）にエディタへ反映する。
  $effect(() => {
    if (editor && value !== editor.getValue()) {
      applyingProp = true
      editor.setValue(value)
      applyingProp = false
    }
  })

  onDestroy(() => editor?.dispose())
</script>

<div class="editor" bind:this={container}></div>

<style>
  .editor {
    height: 100%;
    width: 100%;
  }
</style>
