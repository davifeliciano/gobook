// lissajous generates GIF animations of
// random lissajous figures
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

var palette = []color.Color{color.Black, color.RGBA{26, 255, 0, 1}}

const (
	blackIndex = 0 // First color of the palette
	colorIndex = 1 // Second color of the palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of oscillations on x
		res     = 0.001 // Angular resolution
		size    = 100   // Canvas range is [-size, size]
		nframes = 64    // Number of frames
		delay   = 8     // Frame duration in cs
	)
	freq := rand.Float64() * 3.0 // Relative freq of the y oscillator
	phase := 0.0                 // Phase shift
	anim := gif.GIF{LoopCount: nframes}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < 2*cycles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(freq*t + phase)
			img.SetColorIndex(size+int(size*x+0.5), size+int(size*y+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	// WARNING: Ignoring possible codification errors
}
