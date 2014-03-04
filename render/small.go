package render

import (
	"fmt"
	"image"
	"bytes"

	"github.com/ikaros/tir/colors"
	"github.com/ikaros/tir/offset"
)

func Small(img image.Image) {

	var b bytes.Buffer

	bounds := img.Bounds()

	b.WriteString(offset.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {

		b.WriteString(offset.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			fg(colors.XTerm256.Index(img.At(x, y)), &b)

			if y != bounds.Max.Y - 1 {
				bg(colors.XTerm256.Index(img.At(x, y+1)), &b)
			}

			b.WriteString("▀")
		}

		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	fmt.Print(b.String())
}
