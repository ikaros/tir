package render

import (
	"fmt"
	"image"

	"github.com/ikaros/tir/colors"
	"github.com/ikaros/tir/offset"
)

func Big(img image.Image) {
	var colorIndex int

	bounds := img.Bounds()

	fmt.Print(offset.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

		fmt.Print(offset.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			colorIndex = colors.XTerm256.Index(img.At(x, y))
			fmt.Print(bg(colorIndex), "  ")
		}

		fmt.Print(colorReset, "\n")
	}
}
