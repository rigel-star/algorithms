package main


import (
	"fmt"
	"math"
	"math/rand"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

/*constant variables*/
const (
	noiseW = 128
	noiseH = 128
)

/*global varibles*/
var (
	noise [noiseW][noiseH]float64
)


func main() {
	generateNoise()

	img := image.NewRGBA( image.Rect( 0, 0, noiseW, noiseH ))

	for x := 0; x < noiseW; x++ {
		for y := 0; y < noiseH; y++ {

			//for non smoothed version
			//rgb := uint8( 256 * noise[x / 8.0][y / 8.0] )

			//for smoothed version
			//rgb := uint8( 256 * smoothNoise( float64(x) / 8.0, float64(y) / 8.0) )

			//for smoothed and cluody effect
			rgb := uint8( turbulence( float64(x), float64(y), 256 ))

			col := color.RGBA{R: rgb, G: rgb, B: rgb}
			img.Set( x, y, col )
		}
	}

	saveImg( img, "C:\\Users\\Ramesh\\Desktop\\moreCloudy.jpg" )
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
Generates for cloudy effect.
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
		fmt.Println("Error generated!")
	}

	err = sv.Close()
	if err != nil {
		fmt.Println("Error generated!")
	}
}