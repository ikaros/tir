package render

import (
	"bytes"
	"image"
	"strconv"
)

const colorReset = "\033[0m"

var List []func(image.Image) = []func(image.Image){Small, Big}

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
