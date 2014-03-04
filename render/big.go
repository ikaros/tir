package render

import (
	"bytes"
	"fmt"
	"image"

	"github.com/ikaros/tir/colors"
	"github.com/ikaros/tir/offset"
)

func Big(img image.Image) {
	var b bytes.Buffer

	bounds := img.Bounds()

	b.WriteString(offset.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

		b.WriteString(offset.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			bg(colors.XTerm256.Index(img.At(x, y)), &b)
			b.WriteString("  ")
		}

		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	fmt.Print(b.String())
}
