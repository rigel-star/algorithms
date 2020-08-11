package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)


const (
	w = 400
	h = 300
)


func generate() {
	pImag := 0.0
	pReal := 0.0
	newRe := 0.0
	newIm := 0.0
	oldRe := 0.0
	oldIm := 0.0
	zoom := 1.0
	moveX := -0.5
	moveY := 0.0
	maxIter := 300

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pReal = 1.5 * (float64(x) - w / 2) / (0.5 * zoom * w) + moveX
			pImag = float64(y - h / 2) / (0.5 * zoom * h) + moveY

			newRe = 0
			newIm = 0
			oldIm = 0

			i := 0
			for ; i < maxIter; i++ {
				oldRe = newRe
				oldIm = newIm

				newRe = oldRe * oldRe - oldIm * oldIm + pReal
				newIm = 2 * oldRe * oldIm + pImag

				if (newRe * newRe + newIm * newIm) > 4 {
					break
				}
			}

			col := color.RGBA{uint8(i % 256), 20, 20, 255}
			img.Set(x, y, col)
		}
	}

	saveImg(img, "C:\\Users\\Ramesh\\Desktop\\mset.jpg")
}


func saveImg(img image.Image, path string) {
	sv, err := os.Create(path)
	if err != nil {
		fmt.Println("Error generated!")
	}

	err = jpeg.Encode(sv, img, nil)
	if err != nil {
		fmt.Println("Error generated!")
	}

	err = sv.Close()
	if err != nil {
		fmt.Println("Error generated!")
	}
}


func main() {
	generate()
}
