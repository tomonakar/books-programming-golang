package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x70, 0xD0, 0x88, 0xFF}, color.RGBA{0xD0, 0x80, 0xB8, 0xFF}, color.RGBA{0x70, 0x88, 0xD0, 0xFF}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
	bruIndex   = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	lissajous(os.Stdout)
}

func randColor(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()*(max-min) - min
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンパスは[-size..+size] の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム遅延
	)

	freq := rand.Float64() * 3.0 //発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意：エンコードエラーを無視

}
