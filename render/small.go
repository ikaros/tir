package render

import (
	"fmt"
	"image"

	"github.com/ikaros/tir/colors"
	"github.com/ikaros/tir/offset"
)

func Small(img image.Image) {

	bounds := img.Bounds()

	fmt.Print(offset.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {

		fmt.Print(offset.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			fmt.Print(fg(colors.XTerm256.Index(img.At(x, y))))

			if y != bounds.Max.Y - 1 {
				fmt.Print(bg(colors.XTerm256.Index(img.At(x, y+1))))
			}

			fmt.Print("▀")
		}

		fmt.Print(colorReset, "\n")
	}
}
