package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

func draw(imgW int, imgH int, xoff float64, yoff float64, radius float64) {

	img := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	col := color.RGBA{R: 255, A: 0}

	for theta := 0; theta < 360; theta++ {
		x := radius * math.Cos(float64(theta)) + xoff
		y := radius * math.Sin(float64(theta)) + yoff
		img.Set(int(x), int(y), col)
	}

	save(img, "C:\\Users\\Ramesh\\Desktop\\circle.jpg")
}


func save(img image.Image, path string) {
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
	draw(500, 500, 250, 250, 100)
}
