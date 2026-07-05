<script>
  // 左ペイン: カテゴリー → 題 のツリー。カテゴリーはアコーディオンで開閉できる。
  let { categories = [], currentPath = null, onselect } = $props()

  // 開閉状態。opened[カテゴリー名] === true で開いている（既定は閉じ）。
  let opened = $state({})

  function toggle(name) {
    opened = { ...opened, [name]: !opened[name] }
  }
</script>

<nav class="tree">
  <div class="tree-head">レッスン</div>
  {#each categories as cat (cat.name)}
    <button
      class="category"
      onclick={() => toggle(cat.name)}
      aria-expanded={!!opened[cat.name]}
      title={opened[cat.name] ? '閉じる' : '開く'}
    >
      <span class="chev" class:closed={!opened[cat.name]}>▾</span>
      <span class="cat-name">{cat.name}</span>
    </button>
    {#if opened[cat.name]}
      {#each cat.lessons as lesson (lesson.path)}
        <button
          class="lesson"
          class:active={lesson.path === currentPath}
          onclick={() => onselect(lesson.path)}
        >
          {lesson.title}
        </button>
      {/each}
    {/if}
  {/each}
</nav>

<style>
  .tree {
    height: 100%;
    overflow-y: auto;
    background: #252526;
    padding: 8px 0;
    user-select: none;
  }
  .tree-head {
    font-size: 11px;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: #9ca3af;
    padding: 4px 12px 8px;
  }
  .category {
    display: flex;
    align-items: center;
    gap: 6px;
    width: 100%;
    text-align: left;
    background: none;
    border: none;
    color: #cbd5e1;
    font-size: 12px;
    font-weight: 600;
    padding: 8px 12px 4px;
    cursor: pointer;
  }
  .category:hover {
    color: #fff;
  }
  .chev {
    display: inline-block;
    font-size: 10px;
    color: #9ca3af;
    transition: transform 0.12s ease;
  }
  .chev.closed {
    transform: rotate(-90deg);
  }
  .lesson {
    display: block;
    width: 100%;
    text-align: left;
    background: none;
    border: none;
    color: #d4d4d4;
    padding: 6px 12px 6px 30px;
    font-size: 13px;
    cursor: pointer;
  }
  .lesson:hover {
    background: #2a2d2e;
  }
  .lesson.active {
    background: #094771;
    color: #fff;
  }
</style>
