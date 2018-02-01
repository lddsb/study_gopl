package main

import (
	"image/color"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"net/http"
	"log"
	"strconv"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0x80, 0x00, 0xff}}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		//  parse string to float
		cycles, _ := strconv.ParseFloat(r.Form.Get("cycles"), 64)
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 //  angular resolution
		size    = 100   //  image canvas covers [-size..+size]
		nframes = 64    //  number of animation frames
		delay   = 8     //  delay between frames in 10ms units
	)
	colorIndex := 1
	freq := rand.Float64() * 3.0	//  relative frequency of y oscillator
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0  //  phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(size*x+0.5), size+(int(size*y+0.5)), uint8(colorIndex))
		}
		colorIndex++
		if colorIndex >= len(palette) {
			colorIndex = 1
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
