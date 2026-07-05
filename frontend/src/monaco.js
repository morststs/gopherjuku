// Monaco Editor のスリム構成。
// 完全な言語サービスは使わず、エディタ本体 + Go の構文ハイライト（Monarch）+
// エディタ Worker のみを設定する。
import * as monaco from 'monaco-editor'
import EditorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'
import { format as formatBackend } from './wails.js'

self.MonacoEnvironment = {
  getWorker() {
    return new EditorWorker()
  },
}

// Go 用の整形プロバイダを登録する。これにより Monaco 既定の
// 「Format Document」（Shift+Alt+F）がバックエンドの gofmt を呼ぶ。
// 整形できない（構文エラー等）場合は編集なし（[]）で何もしない。
monaco.languages.registerDocumentFormattingEditProvider('go', {
  displayName: 'gofmt',
  async provideDocumentFormattingEdits(model) {
    const res = await formatBackend(model.getValue())
    if (!res?.success) return []
    return [{ range: model.getFullModelRange(), text: res.formatted }]
  },
})

export { monaco }
