package render

import (
	"bytes"
	"fmt"
	"image"
	"image/color"

	"github.com/ikaros/tir/colors"
	"github.com/ikaros/tir/offset"
)

func Big(img image.Image) {
	var(
		b bytes.Buffer
		currentColor color.Color
	)

	bounds := img.Bounds()

	b.WriteString(offset.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

		b.WriteString(offset.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			currentColor = img.At(x, y)

			if isTransparent(currentColor) {
				b.WriteString(colorReset)
			}else {
				bg(colors.XTerm256.Index(currentColor), &b)
			}

			b.WriteString("  ")
		}

		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	fmt.Print(b.String())
}
