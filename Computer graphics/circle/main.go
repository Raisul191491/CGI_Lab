package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	width  = 200
	height = 200
)

func setPixel(img *image.RGBA, x, y int, c color.Color) {
	img.Set(x, y, c)
}

func drawCircleBresenham(img *image.RGBA, centerX, centerY, radius int, col color.Color) {
	x := 0
	y := radius
	d := 3 - 2*radius

	for x <= y {
		setPixel(img, centerX+x, centerY+y, col)
		setPixel(img, centerX+y, centerY+x, col)
		setPixel(img, centerX-y, centerY+x, col)
		setPixel(img, centerX-x, centerY+y, col)
		setPixel(img, centerX-x, centerY-y, col)
		setPixel(img, centerX-y, centerY-x, col)
		setPixel(img, centerX+y, centerY-x, col)
		setPixel(img, centerX+x, centerY-y, col)

		if d <= 0 {
			d = d + (4 * x) + 6
		} else {
			d = d + (4 * x) - (4 * y) + 10
			y--
		}
		x++
	}
}

func drawCircleMidpoint(img *image.RGBA, centerX, centerY, radius int, col color.Color) {
	x := 0
	y := radius
	d := 1 - radius

	for x <= y {
		setPixel(img, centerX+x, centerY+y, col)
		setPixel(img, centerX+y, centerY+x, col)
		setPixel(img, centerX-y, centerY+x, col)
		setPixel(img, centerX-x, centerY+y, col)
		setPixel(img, centerX-x, centerY-y, col)
		setPixel(img, centerX-y, centerY-x, col)
		setPixel(img, centerX+y, centerY-x, col)
		setPixel(img, centerX+x, centerY-y, col)

		if d <= 0 {
			d = d + (2 * x) + 3
		} else {
			d = d + (2 * x) - (2 * y) + 5
			y--
		}
		x++
	}
}

func main() {
	img1 := image.NewRGBA(image.Rect(0, 0, width, height))
	img2 := image.NewRGBA(image.Rect(0, 0, width, height))

	drawCircleBresenham(img1, 100, 100, 60, color.White)
	drawCircleBresenham(img2, 100, 100, 60, color.White)

	bFile, err := os.Create("circleBresenham.png")
	if err != nil {
		panic(err)
	}
	defer bFile.Close()

	err = png.Encode(bFile, img1)
	if err != nil {
		panic(err)
	}

	mFile, err := os.Create("circleMidpoint.png")
	if err != nil {
		panic(err)
	}
	defer bFile.Close()

	err = png.Encode(mFile, img2)
	if err != nil {
		panic(err)
	}
}
