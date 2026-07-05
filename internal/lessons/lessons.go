// Package lessons は、埋め込み FS 上の教材（_contents）を読み出して
// 左ペインのツリー（カテゴリー → 題）を組み立てる。
//
// 表示順は「数字プレフィックス」で制御する。ディレクトリ名・ファイル名の先頭に
// 連番を付け（例: `1_基本/`, `01 ハローワールド.go`）、その数値で並べ替える。
// 表示名・ファイル探索用パスは以下のように扱う:
//   - 表示名（Category.Name / Lesson.Title）… プレフィックスを除去
//   - Lesson.Path … プレフィックス込みの実ファイルパス（Source の探索に使う）
//
// Wails 非依存なので fstest.MapFS で単体テストできる（lessons_test.go）。
package lessons

import (
	"io/fs"
	"path"
	"sort"
	"strconv"
	"strings"
)

// contentsRoot は埋め込み FS 上の教材ルート。ビルド対象から外すため "_" 始まり。
const contentsRoot = "_contents"

// Lesson は 1 件の教材（コードの題）。
type Lesson struct {
	Category string `json:"category"` // 表示用カテゴリー名（例: "基本"）
	Title    string `json:"title"`    // 表示用の題（例: "ハローワールド"）
	Path     string `json:"path"`     // 実ファイルパス（例: "1_基本/01 ハローワールド"・拡張子なし）
}

// Category は左ペインの 1 カテゴリーとその配下のレッスン。
type Category struct {
	Name    string   `json:"name"`
	Lessons []Lesson `json:"lessons"`
}

// Service は教材リポジトリ。
type Service struct {
	fsys fs.FS
}

func New(fsys fs.FS) *Service { return &Service{fsys: fsys} }

// Tree は「カテゴリー → 題」の一覧を数字プレフィックス順で返す。
func (s *Service) Tree() []Category {
	entries, err := fs.ReadDir(s.fsys, contentsRoot)
	if err != nil {
		return []Category{}
	}

	type ordered struct {
		ord int
		cat Category
	}
	var cats []ordered

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		dirName := e.Name()
		catOrd, catName := parseOrder(dirName)

		files, err := fs.ReadDir(s.fsys, path.Join(contentsRoot, dirName))
		if err != nil {
			continue
		}

		type olesson struct {
			ord    int
			lesson Lesson
		}
		var lessons []olesson
		for _, f := range files {
			if f.IsDir() || !strings.HasSuffix(f.Name(), ".go") {
				continue
			}
			base := strings.TrimSuffix(f.Name(), ".go")
			lessonOrd, title := parseOrder(base)
			lessons = append(lessons, olesson{
				ord: lessonOrd,
				lesson: Lesson{
					Category: catName,
					Title:    title,
					Path:     dirName + "/" + base,
				},
			})
		}
		if len(lessons) == 0 {
			continue
		}
		sort.SliceStable(lessons, func(i, j int) bool { return lessons[i].ord < lessons[j].ord })

		out := make([]Lesson, len(lessons))
		for i := range lessons {
			out[i] = lessons[i].lesson
		}
		cats = append(cats, ordered{ord: catOrd, cat: Category{Name: catName, Lessons: out}})
	}

	sort.SliceStable(cats, func(i, j int) bool { return cats[i].ord < cats[j].ord })
	result := make([]Category, len(cats))
	for i := range cats {
		result[i] = cats[i].cat
	}
	return result
}

// Source は指定レッスンの Go ソースを返す。p は Tree が返す Path
//（例 "1_基本/01 ハローワールド"）。パストラバーサルは拒否する。
func (s *Service) Source(p string) (string, error) {
	clean := path.Clean("/" + p) // 先頭 ".." を無害化
	rel := strings.TrimPrefix(clean, "/")
	b, err := fs.ReadFile(s.fsys, path.Join(contentsRoot, rel+".go"))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// parseOrder は "01 ハローワールド" や "1_基本" のような名前を、先頭の連番と
// それ以降（区切りのスペース/アンダースコアを除いた表示名）に分解する。
// 連番が無ければ末尾扱い（大きな値）とし、名前はそのまま返す。
func parseOrder(name string) (int, string) {
	i := 0
	for i < len(name) && name[i] >= '0' && name[i] <= '9' {
		i++
	}
	if i == 0 {
		return 1 << 30, name
	}
	n, _ := strconv.Atoi(name[:i])
	rest := strings.TrimLeft(name[i:], " _")
	return n, rest
}
