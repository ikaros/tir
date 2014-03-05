package render

import (
	"bytes"
	"fmt"
	"image"
	"image/color"

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
			pixel(&b, img.At(x, y), img.At(x, y+1))
		}

		b.WriteString(colorReset)
		b.WriteString("\n")
	}

	fmt.Print(b.String())
}

func pixel(b *bytes.Buffer, top color.Color, bottom color.Color) {
	b.WriteString(colorReset)

	if isTransparent(top) && isTransparent(bottom) {
		b.WriteString(" ")

	} else if isTransparent(top) {
		fg(colors.XTerm256.Index(bottom), b)
		b.WriteString("▄")

	} else if isTransparent(bottom) {
		fg(colors.XTerm256.Index(top), b)
		b.WriteString("▀")

	} else {
		fg(colors.XTerm256.Index(top), b)
		bg(colors.XTerm256.Index(bottom), b)
		b.WriteString("▀")
	}
}
