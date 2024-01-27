## Bresenham Line Drawing Algorithm 
---

```golang

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	width  = 1000
	height = 400
)

func setPixel(img *image.RGBA, x, y int, c color.Color) {
	img.Set(x, y, c)
}

func drawLine(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	var dx, dy int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, abs(y2-y1)

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		setPixel(img, x1, y1, col)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			setPixel(img, x1, y1, col)
			x1++
		}
		setPixel(img, x1, y1, col)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			setPixel(img, x1, y1, col)
			y1++
		}
		setPixel(img, x1, y1, col)

	// Is line a diagonal ?
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

	// wider than high ?
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

		// higher than wide.
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

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw "S"
	drawLine(img, 200, 100, 400, 100, color.White)
	drawLine(img, 200, 100, 200, 200, color.White)
	drawLine(img, 200, 200, 400, 200, color.White)
	drawLine(img, 400, 200, 400, 300, color.White)
	drawLine(img, 200, 300, 400, 300, color.White)

	// Draw "W"
	drawLine(img, 500, 100, 550, 300, color.White)
	drawLine(img, 550, 300, 600, 100, color.White)
	drawLine(img, 600, 100, 650, 300, color.White)
	drawLine(img, 650, 300, 700, 100, color.White)

	// Draw "E"
	drawLine(img, 750, 100, 750, 300, color.White)
	drawLine(img, 750, 100, 900, 100, color.White)
	drawLine(img, 750, 200, 900, 200, color.White)
	drawLine(img, 750, 300, 900, 300, color.White)

	// Save the image to a file
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func abs(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}

```


## DDA Line Drawing Algorithm
---

```golang

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
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw "S"
	drawLineDDA(img, 200, 100, 400, 100, color.White)
	drawLineDDA(img, 200, 100, 200, 200, color.White)
	drawLineDDA(img, 200, 200, 400, 200, color.White)
	drawLineDDA(img, 400, 200, 400, 300, color.White)
	drawLineDDA(img, 200, 300, 400, 300, color.White)

	// Draw "W"
	drawLineDDA(img, 500, 100, 550, 300, color.White)
	drawLineDDA(img, 550, 300, 600, 100, color.White)
	drawLineDDA(img, 600, 100, 650, 300, color.White)
	drawLineDDA(img, 650, 300, 700, 100, color.White)

	// Draw "E"
	drawLineDDA(img, 750, 100, 750, 300, color.White)
	drawLineDDA(img, 750, 100, 900, 100, color.White)
	drawLineDDA(img, 750, 200, 900, 200, color.White)
	drawLineDDA(img, 750, 300, 900, 300, color.White)

	// Save the image to a file
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
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