package render

import (
	"bytes"
	"image"
	"image/color"
	"strconv"
)

const colorReset = "\033[0m"

var List []func(image.Image) = []func(image.Image){Small, Big}

func isTransparent(c color.Color) bool {
	_, _, _, a := c.RGBA()
	return a == 0
}

func lastRow(b image.Rectangle, i int) bool {
	return i == b.Max.Y - 1
}

func fg(colorCode int, b *bytes.Buffer) {
	b.WriteString("\033[38;5;")
	b.WriteString(strconv.Itoa(colorCode))
	b.WriteString("m")
}

func bg(colorCode int, b *bytes.Buffer) {
	b.WriteString("\033[48;5;")
	b.WriteString(strconv.Itoa(colorCode))
	b.WriteString("m")
}
