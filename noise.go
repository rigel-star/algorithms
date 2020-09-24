package main

import (
	"colorica"
	"fmt"
	"math"
	"math/rand"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const (
	noiseW = 128
	noiseH = 128
)

var (
	noise [noiseW][noiseH]float64
	savePath string = "C:\\Users\\Ramesh\\Desktop\\moreCloudy.jpg"
)


func main() {
	generateNoise()
	//saveImg( generateWoodNoise(), savePath )

	img := image.NewRGBA( image.Rect( 0, 0, noiseW, noiseH ))

	var L float64
	var torgb colorica.RGB
	var hsl colorica.HSL
	var rgb color.RGBA

	for x := 0; x < noiseW; x++ {
		for y := 0; y < noiseH; y++ {

			L = 192 + turbulence( float64(x), float64(y), 64 ) / 4
			hsl = colorica.HSL{169.0, 255.0, L} //H = 169, S = 255, L = L
			torgb = hsl.ToRGB()

			fr := torgb.R * 255.0
			fg := torgb.G * 255.0
			fb := torgb.B * 255.0

			//for colored pattern
			rgb = color.RGBA{R: uint8(fr), G: uint8(fg), B: uint8(fb)}

			//for non smoothed version
			//rgb := uint8( noise[x >> 3][y >> 3] << 8 )

			//for smoothed version
			//rgb := uint8( smoothNoise( float64(x) / 8.0, float64(y) / 8.0) << 8 )

			//for smoothed and cluody effect
			//rgb := uint8( turbulence( float64(x), float64(y), 256 ))

			//col := color.RGBA{R: rgb, G: rgb, B: rgb}
			img.Set( x, y, rgb )
		}
	}

	saveImg( img, savePath )
}


func generateWoodNoise() image.Image {
	xyPeriod := 12.0 //number of rings
  	turbPower := 0.1 //makes twists
 	turbSize := 32.0 //initial size of the turbulence

 	img := image.NewRGBA( image.Rect( 0, 0, noiseW, noiseH ))

 	for x := 0; x < noiseW; x++ {
		for y := 0; y < noiseH; y++ {

			xValue := (float64(x) - noiseW / 2.0) / noiseW
		    yValue := (float64(y) - noiseH / 2.0) / noiseW
		    distValue := math.Sqrt( xValue * xValue + yValue * yValue ) +
		    				turbPower * turbulence( float64(x), float64(y), turbSize ) / 256.0
		    sineValue := 128.0 * math.Abs( math.Sin(2.0 * xyPeriod * distValue * 3.14159 ));
		    col := color.RGBA{R:uint8( 80 + sineValue ), 
								G:uint8( 30 + sineValue ), B:30}
			img.Set( x, y, col )
		}
	}
	return img
}


/**
Generate random noise with the help of random function.
This generates the random noise value between 0 & 32767 and
mods with 32768 and returns the value dividing by 32768 to 
keep it in range of 0 - 1.
*/
func generateNoise() {
	for x := 0; x < noiseW; x++ {
		for y := 0; y < noiseH; y++ {
			noise[x][y] = ( math.Mod( random( 0, 32767 ), 32768 ) ) / 32768.0
		}
	}
}


/**
Smooth outs the generated noise.
*/
func smoothNoise( x, y float64 ) float64 {
	//get fractional part of x and y
	//for eg. 1.2 - 1 = .2
   	fractX := x - float64( int(x) )
   	fractY := y - float64( int(y) )

   	//wrap around
   	x1 := ( int(x) + noiseW ) % noiseW
   	y1 := ( int(y) + noiseH ) % noiseH

   	x2 := ( x1 + noiseW - 1 ) % noiseW
   	y2 := ( y1 + noiseH - 1 ) % noiseH

   	value := 0.0
   	value += fractX * fractY * noise[x1][y1]
   	value += ( 1 - fractX ) * fractY * noise[x1][y2];
   	value += fractX * ( 1 - fractY ) * noise[x2][y1];
   	value += ( 1 - fractX ) * ( 1 - fractY ) * noise[x2][y2];

   	return value
}


/**
Generates cloudy effect.
*/
func turbulence( x, y, size float64 ) float64 {
	value := 0.0
	startingSize := size

	for size >= 1 {
		value += smoothNoise( x / size, y / size) * size
		size /= 2.0
	}

	return ( 128.0 * value / startingSize )
}


/**
Random value between specified min and max constrain.
*/
func random( min, max int ) float64 {
    return float64( rand.Intn( max - min ) + min )
}


func saveImg( img image.Image, path string ) {
	sv, err := os.Create( path )
	if err != nil {
		fmt.Println( "Error generated!" )
	}

	err = jpeg.Encode( sv, img, nil )
	if err != nil {
		fmt.Println( "Error generated!" )
	}

	err = sv.Close()
	if err != nil {
		fmt.Println( "Error generated!" )
	}
}