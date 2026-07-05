package lessons

import (
	"testing"
	"testing/fstest"
)

func TestTreeOrderingAndDisplay(t *testing.T) {
	// わざと辞書順とは違う（数字順で正しく並ぶべき）構成にする。
	fsys := fstest.MapFS{
		"_contents/2_標準ライブラリ/04 sort.go": {Data: []byte("package main")},
		"_contents/2_標準ライブラリ/02 strings.go": {Data: []byte("package main")},
		"_contents/1_基本/02 変数と型.go":     {Data: []byte("package main")},
		"_contents/1_基本/01 ハローワールド.go": {Data: []byte("package main")},
		"_contents/3_応用/08 ジェネリクス.go":  {Data: []byte("package main")},
	}
	cats := New(fsys).Tree()

	if len(cats) != 3 {
		t.Fatalf("カテゴリ数=%d want 3", len(cats))
	}
	// カテゴリーは数字プレフィックス順・表示名はプレフィックス除去。
	wantCats := []string{"基本", "標準ライブラリ", "応用"}
	for i, w := range wantCats {
		if cats[i].Name != w {
			t.Errorf("cat[%d]=%q want %q", i, cats[i].Name, w)
		}
	}
	// 基本の題は 01→02 の順、表示名はプレフィックス除去。
	b := cats[0].Lessons
	if len(b) != 2 || b[0].Title != "ハローワールド" || b[1].Title != "変数と型" {
		t.Fatalf("基本の題順/表示が不正: %+v", b)
	}
	// Path は探索用にプレフィックス保持。
	if b[0].Path != "1_基本/01 ハローワールド" {
		t.Errorf("path=%q want %q", b[0].Path, "1_基本/01 ハローワールド")
	}
}

func TestSourceReadsByPath(t *testing.T) {
	fsys := fstest.MapFS{
		"_contents/1_基本/01 ハローワールド.go": {Data: []byte("package main // hi")},
	}
	src, err := New(fsys).Source("1_基本/01 ハローワールド")
	if err != nil {
		t.Fatal(err)
	}
	if src != "package main // hi" {
		t.Errorf("src=%q", src)
	}
}

func TestParseOrder(t *testing.T) {
	cases := []struct {
		in    string
		ord   int
		title string
	}{
		{"01 ハローワールド", 1, "ハローワールド"},
		{"1_基本", 1, "基本"},
		{"12 インターフェース", 12, "インターフェース"},
		{"プレフィックス無し", 1 << 30, "プレフィックス無し"},
	}
	for _, c := range cases {
		o, tt := parseOrder(c.in)
		if o != c.ord || tt != c.title {
			t.Errorf("parseOrder(%q)=(%d,%q) want (%d,%q)", c.in, o, tt, c.ord, c.title)
		}
	}
}
