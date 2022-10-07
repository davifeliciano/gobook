// lissajousserver generates randnom lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // First color of the palette
	blackIndex = 1 // Second color of the palette
)

func main() {
	const res float64 = 0.001
	handler := func(w http.ResponseWriter, r *http.Request) {
		argsStringValues := map[string]string{"cycles": "5", "size": "100", "nframes": "64", "delay": "8"}
		argsIntValues := make(map[string]int)
		for k := range argsStringValues {
			if queryValue := r.URL.Query().Get(k); queryValue != "" {
				argsStringValues[k] = queryValue
			}
		}
		for k, v := range argsStringValues {
			if intValue, err := strconv.Atoi(v); err != nil {
				log.Fatal(err)
			} else {
				argsIntValues[k] = intValue
			}
		}
		cycles := argsIntValues["cycles"]
		size := argsIntValues["size"]
		nframes := argsIntValues["nframes"]
		delay := argsIntValues["delay"]
		lissajous(w, cycles, res, size, nframes, delay)
	}
	http.HandleFunc("/", handler) // each request of route / will call the func handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	// cycles is the number of oscillations on x
	// res is the angular resolution
	// canvas range is [-size, size]
	// nframes is the number of frames
	// delay is the frame duration in cs
	freq := rand.Float64() * 3.0 // Relative freq of the y oscillator
	phase := 0.0                 // Phase shift
	anim := gif.GIF{LoopCount: nframes}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(2*cycles)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(freq*t + phase)
			img.SetColorIndex(
				size+int(float64(size)*x+0.5),
				size+int(float64(size)*y+0.5),
				blackIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	// WARNING: Ignoring possible codification errors
}
