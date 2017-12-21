package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func drawMandelbrotImg2File() {

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("mandelbrot.png")
	if err != nil {
		// Handle error
	}
	img := buildRGBA()
	png.Encode(outputFile, img) // NOTE: ignoring errors
	outputFile.Close()
}

func drawMandelbroImg2HTMLTag() {
	// In-memory buffer to store PNG image
	// before we base 64 encode it
	var buff bytes.Buffer
	img := buildRGBA()

	// The Buffer satisfies the Writer interface so we can use it with Encode
	// In previous example we encoded to a file, this time to a temp buffer
	png.Encode(&buff, img)

	// Encode the bytes in the buffer to a base64 string
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	// You can embed it in an html doc with this string
	htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"
	fmt.Println(htmlImage)
}

func buildRGBA() *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func main() {
	drawMandelbrotImg2File()   //第一种方式呈现：输出到文件
	drawMandelbroImg2HTMLTag() //第二种方式呈现：生成HTML图形标签，直接在html文件中嵌入这行代码即可在浏览器中看到图片。
}
