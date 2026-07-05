// gen-lessons は _contents/ の教材を、ブラウザ版（GitHub Pages）が読み込む
// 静的 JSON に書き出すツール。バックエンドの無い Pages 版では、この JSON から
// レッスンツリーとソースを供給する。
//
// 使い方: go run ./cmd/gen-lessons <出力先.json>（リポジトリルートで実行）
package main

import (
	"encoding/json"
	"log"
	"os"

	"gopherjuku/internal/lessons"
)

type webLesson struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	Path     string `json:"path"`
	Source   string `json:"source"`
}

type webCategory struct {
	Name    string      `json:"name"`
	Lessons []webLesson `json:"lessons"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("使い方: gen-lessons <出力先.json>")
	}
	out := os.Args[1]

	// リポジトリルートの _contents/ を読む（internal/lessons と同じロジック）。
	svc := lessons.New(os.DirFS("."))

	var cats []webCategory
	for _, c := range svc.Tree() {
		wc := webCategory{Name: c.Name}
		for _, l := range c.Lessons {
			src, err := svc.Source(l.Path)
			if err != nil {
				log.Fatalf("ソース読み込み失敗 %q: %v", l.Path, err)
			}
			wc.Lessons = append(wc.Lessons, webLesson{
				Category: l.Category,
				Title:    l.Title,
				Path:     l.Path,
				Source:   src,
			})
		}
		cats = append(cats, wc)
	}

	b, err := json.MarshalIndent(map[string]any{"categories": cats}, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(out, b, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s (%d カテゴリー)", out, len(cats))
}
