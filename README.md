# Computer Graphics Lab

## Bresenham & DDA Line Drawing Algorithm 
---
```go
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	width  = 1000
	height = 400
)

func setPixel(img *image.RGBA, x, y int, c color.Color) {
	img.Set(x, y, c)
}

func drawLineBresenham(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	var dx, dy int

	// Drawing p1 -> p2 == Drawing p2 -> p1,
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, abs(y2-y1)

	switch {

	// For point
	case x1 == x2 && y1 == y2:
		setPixel(img, x1, y1, col)

	// Horizontal line
	case y1 == y2:
		for ; dx != 0; dx-- {
			setPixel(img, x1, y1, col)
			x1++
		}
		setPixel(img, x1, y1, col)

	// Vertical line
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			setPixel(img, x1, y1, col)
			y1++
		}
		setPixel(img, x2, y2, col)

	// Diagonal m == 1
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				setPixel(img, x1, y1, col)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				setPixel(img, x1, y1, col)
				x1++
				y1--
			}
		}
		setPixel(img, x1, y1, col)

	// Width > Height
	case dx > dy:
		p := 2*dy - dx
		for x1 < x2 {
			if p >= 0 {
				setPixel(img, x1, y1, col)
				y1++
				p = p + 2*dy - 2*dx
			} else {
				setPixel(img, x1, y1, col)
				p = p + 2*dy
			}
			x1++
		}
		setPixel(img, x2, y2, col)

	// Height > Width
	case dx < dy:
		p := 2*dx - dy
		if y1 < y2 {
			for y1 < y2 {
				setPixel(img, x1, y1, col)
				if p >= 0 {
					x1++
					p = p + 2*dx - 2*dy
				} else {
					p = p + 2*dx
				}
				y1++
			}
			setPixel(img, x2, y2, col)
		} else {
			for y1 > y2 {
				setPixel(img, x1, y1, col)
				if p >= 0 {
					x1++
					p = p + 2*dx - 2*dy
				} else {
					p = p + 2*dx
				}
				y1--
			}
			setPixel(img, x2, y2, col)
		}

	}
}

func drawLineDDA(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	dx := x2 - x1
	dy := y2 - y1

	var steps int

	if abs(dx) > abs(dy) {
		steps = abs(dx)
	} else {
		steps = abs(dy)
	}

	xIncrement := float64(dx) / float64(steps)
	yIncrement := float64(dy) / float64(steps)

	x, y := float64(x1), float64(y1)

	for i := 0; i < steps; i++ {
		setPixel(img, int(math.Round(x)), int(math.Round(y)), col)
		x += xIncrement
		y += yIncrement
	}

	setPixel(img, x2, y2, col)
}

func main() {
	bresenham := image.NewRGBA(image.Rect(0, 0, width, height))
	dda := image.NewRGBA(image.Rect(0, 0, width, height))

	// Bresenham's algorithm
	// Draw "S"
	drawLineBresenham(bresenham, 200, 100, 400, 100, color.White)
	drawLineBresenham(bresenham, 200, 100, 200, 200, color.White)
	drawLineBresenham(bresenham, 200, 200, 400, 200, color.White)
	drawLineBresenham(bresenham, 400, 200, 400, 300, color.White)
	drawLineBresenham(bresenham, 200, 300, 400, 300, color.White)

	// Draw "W"
	drawLineBresenham(bresenham, 500, 100, 550, 300, color.White)
	drawLineBresenham(bresenham, 550, 300, 600, 100, color.White)
	drawLineBresenham(bresenham, 600, 100, 650, 300, color.White)
	drawLineBresenham(bresenham, 650, 300, 700, 100, color.White)

	// Draw "E"
	drawLineBresenham(bresenham, 750, 100, 750, 300, color.White)
	drawLineBresenham(bresenham, 750, 100, 900, 100, color.White)
	drawLineBresenham(bresenham, 750, 200, 900, 200, color.White)
	drawLineBresenham(bresenham, 750, 300, 900, 300, color.White)

	// DDA
	// Draw "S"
	drawLineDDA(dda, 200, 100, 400, 100, color.White)
	drawLineDDA(dda, 200, 100, 200, 200, color.White)
	drawLineDDA(dda, 200, 200, 400, 200, color.White)
	drawLineDDA(dda, 400, 200, 400, 300, color.White)
	drawLineDDA(dda, 200, 300, 400, 300, color.White)

	// Draw "W"
	drawLineDDA(dda, 500, 100, 550, 300, color.White)
	drawLineDDA(dda, 550, 300, 600, 100, color.White)
	drawLineDDA(dda, 600, 100, 650, 300, color.White)
	drawLineDDA(dda, 650, 300, 700, 100, color.White)

	// Draw "E"
	drawLineDDA(dda, 750, 100, 750, 300, color.White)
	drawLineDDA(dda, 750, 100, 900, 100, color.White)
	drawLineDDA(dda, 750, 200, 900, 200, color.White)
	drawLineDDA(dda, 750, 300, 900, 300, color.White)

	bFile, err := os.Create("bresenham.png")
	if err != nil {
		panic(err)
	}
	defer bFile.Close()

	err = png.Encode(bFile, bresenham)
	if err != nil {
		panic(err)
	}

	dFile, err := os.Create("dda.png")
	if err != nil {
		panic(err)
	}
	defer dFile.Close()

	err = png.Encode(dFile, dda)
	if err != nil {
		panic(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

```

## Bresenham & Midpoint circle drawing

```golang

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

```
