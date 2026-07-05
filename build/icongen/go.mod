// アイコン生成ツール専用の入れ子モジュール。独自 go.mod を持つことで、
// 親モジュール（gopherjuku）の `go build ./...` / `go vet ./...` の対象外になる。
module icongen

go 1.23.0
