package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

/**
	Color conversion techniques are from : [https://stackoverflow.com/questions/2353211/hsl-to-rgb-color-conversion]
*/

func generateWheel(){

	height := 1024
	width := 1024

	//img := image.Image()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//image center
	centerX := width / 2
	centerY := height / 2
	radius := (width / 2) * (width / 2)

	redX := width
	redY := height / 2
	redRad := width * width

	greenX := 0
	greenY := height / 2
	greenRad := width * width

	blueX := width / 2
	blueY := height
	blueRad := width * width

	for x := 0; x < width ;x++ {
		for y := 0; y < height; y++ {

			a := x - centerX
			b := y - centerY

			dist := a * a + b * b

			if dist < radius {

				rdx := y - redX
				rdy := y - redY
				redDist := rdx * rdx + rdy * rdy
				redVal := 255 - ((float64(redDist) / float64(redRad)) * 256)

				gdx := y - greenX
				gdy := y - greenY
				greenDist := gdx * gdx + gdy * gdy
				greenVal := 255 - ((float64(greenDist) / float64(greenRad)) * 256)

				bdx := y - blueX
				bdy := y - blueY
				blueDist := bdx * bdx + bdy * bdy
				blueVal := 255 - ((float64(blueDist) / float64(blueRad)) * 256)

				col := color.RGBA{uint8(blueVal), uint8(greenVal), uint8(redVal), 255}
				h, s, v := rgbToHSV(float64(col.R), float64(col.G), float64(col.B))

				if v > 1000 {
					fmt.Println("Oh no")
				}

				r, g, b := hsvToRgb(h, s, 1)
				rgb := color.RGBA{uint8(b), uint8(g), uint8(r), 255}
				img.Set(x, y, rgb)

			} else {
				img.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}

	saveImg2(img, "C:\\Users\\Ramesh\\Desktop\\cwheelgo.jpg")
}


func hsvToRgb(h float64, s float64, v float64) (float64, float64, float64){
	var r, g, b float64

	if s == 0 {
		r = v
		g = v
		b = v // achromatic
	} else {
		p, q := 0.0, 0.0

		if v < 0.5 {
			q = v * (1 + s)
		} else {
			q = v + s - v + s
		}

		p = v * (2 - q)

		r = hue2rgb(p, q, h + (1 / 3))
		g = hue2rgb(p, q, h)
		b = hue2rgb(p, q, h - (1 / 3))
	}

	return math.Round(r * 255), math.Round(g * 255), math.Round(b * 255)
}


func hue2rgb(p float64, q float64, t float64) float64 {
	if t < 0.0 {t += 1}
	if t > 1 {t -= 1}
	if t < (1/6) {return p + (q - p) * 6 * t}
	if t < (1/2) {return q}
	if t < (2/3) {return p + (q - p) * (2/3 - t) * 6}
	return p
}


func rgbToHSV(r float64, g float64, b float64) (float64, float64, float64){
	r /= 255
	g /= 255
	b /= 255

	max := math.Max(r, g)
	max = math.Max(max, b)

	min := math.Min(r, g)
	min = math.Min(min, b)

	h := (max + min) / 2
	s := (max + min) / 2
	v := (max + min) / 2

	if max == min {
		h = 0.0
		s = 0.0

	} else {
		d := max - min

		if v > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {

		case r:
			if g < b {
				h = (g - b) / d + 6
				break
			} else {
				h = (g - b) / d
				break
			}

		case g:
			h = (b - r) / (d + 2)
			break

		case b:
			h = (r - g) / (d + 4)
		}

		h /= 6
	}

	return h, s, v
}


func saveImg2(img image.Image, path string) {
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
	generateWheel()
}
