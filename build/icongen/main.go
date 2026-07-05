// Gopherjuku のアプリアイコンを生成する。
// デザイン: Go ブランドのシアン(#00ADD8)の角丸四角 + 白い再生マーク（▶＝実行）。
// 4x4 スーパーサンプリングでエッジを滑らかにする。出力先は引数で指定。
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const S = 1024

// 角丸四角の内側判定（[0,w]x[0,h]、角半径 r）。
func inRoundedRect(x, y, w, h, r float64) bool {
	cx, cy := x, y
	if cx < r {
		cx = r
	} else if cx > w-r {
		cx = w - r
	}
	if cy < r {
		cy = r
	} else if cy > h-r {
		cy = h - r
	}
	dx, dy := x-cx, y-cy
	return dx*dx+dy*dy <= r*r
}

func sign(x1, y1, x2, y2, x3, y3 float64) float64 {
	return (x1-x3)*(y2-y3) - (x2-x3)*(y1-y3)
}

func inTriangle(px, py, ax, ay, bx, by, cx, cy float64) bool {
	d1 := sign(px, py, ax, ay, bx, by)
	d2 := sign(px, py, bx, by, cx, cy)
	d3 := sign(px, py, cx, cy, ax, ay)
	neg := d1 < 0 || d2 < 0 || d3 < 0
	pos := d1 > 0 || d2 > 0 || d3 > 0
	return !(neg && pos)
}

func main() {
	out := "appicon.png"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}

	img := image.NewRGBA(image.Rect(0, 0, S, S))

	// 再生マーク（右向き三角形）を中央やや右に置く。
	ax, ay := 400.0, 312.0
	bx, by := 400.0, 712.0
	cx, cy := 748.0, 512.0

	const N = 4 // スーパーサンプリング数（NxN）
	for y := 0; y < S; y++ {
		for x := 0; x < S; x++ {
			var rSum, gSum, bSum, opaque float64
			for sy := 0; sy < N; sy++ {
				for sx := 0; sx < N; sx++ {
					fx := float64(x) + (float64(sx)+0.5)/N
					fy := float64(y) + (float64(sy)+0.5)/N
					switch {
					case inTriangle(fx, fy, ax, ay, bx, by, cx, cy):
						rSum, gSum, bSum = rSum+255, gSum+255, bSum+255
						opaque++
					case inRoundedRect(fx, fy, S, S, 180):
						gSum, bSum = gSum+173, bSum+216 // R=0
						opaque++
					}
				}
			}
			if opaque == 0 {
				continue // 透明のまま
			}
			alpha := opaque / float64(N*N) * 255
			img.Set(x, y, color.RGBA{
				R: uint8(rSum/opaque + 0.5),
				G: uint8(gSum/opaque + 0.5),
				B: uint8(bSum/opaque + 0.5),
				A: uint8(alpha + 0.5),
			})
		}
	}

	f, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
