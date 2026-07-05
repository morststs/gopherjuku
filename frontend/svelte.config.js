import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'

// flowbite-svelte などが配布する TypeScript 入り .svelte をビルドできるよう
// vitePreprocess を有効化する。Svelte 5 の組み込み TS 除去では一部の型注釈が
// 残って Rollup ビルドが失敗するため、script: true で esbuild による
// TS トランスパイルを明示的に有効化する。この設定を消すとビルドが壊れる。
export default {
  preprocess: vitePreprocess({ script: true }),
}
