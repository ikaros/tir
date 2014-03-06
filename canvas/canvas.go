package canvas

import (
	"bytes"
	"image/color"
	"strconv"

	"github.com/ikaros/tir/colors"
)

type Canvas struct {
	Buffer *bytes.Buffer

	CurrentFG int
	CurrentBG int
}

func New(b *bytes.Buffer) (c Canvas) {
	c.Buffer = b
	return
}

func (c *Canvas) IsTransparent(cl color.Color) bool {
	_, _, _, a := cl.RGBA()
	return a == 0
}

func (c *Canvas) FG(colorCode int) {
	c.Buffer.WriteString("\033[38;5;")
	c.Buffer.WriteString(strconv.Itoa(colorCode))
	c.Buffer.WriteString("m")
}

func (c *Canvas) BG(colorCode int) {
	c.Buffer.WriteString("\033[48;5;")
	c.Buffer.WriteString(strconv.Itoa(colorCode))
	c.Buffer.WriteString("m")
}

func (c *Canvas) TopPixel() {
	c.Buffer.WriteString("▀")
}

func (c *Canvas) BottomPixel() {
	c.Buffer.WriteString("▄")
}

func (c *Canvas) EmptyPixel() {
	c.Buffer.WriteString(" ")
}

func (c *Canvas) Reset() {
	c.Buffer.WriteString("\033[0m")
}

func (c *Canvas) NewLine() {
	c.Buffer.WriteString("\n")
}

func (c *Canvas) Pixel(top color.Color, bottom color.Color) {
	c.Reset()

	if c.IsTransparent(top) && c.IsTransparent(bottom) {
		c.EmptyPixel()

	} else if c.IsTransparent(top) {
		c.FG(colors.XTerm256.Index(bottom))
		c.BottomPixel()

	} else if c.IsTransparent(bottom) {
		c.FG(colors.XTerm256.Index(top))
		c.TopPixel()

	} else {
		c.FG(colors.XTerm256.Index(top))
		c.BG(colors.XTerm256.Index(bottom))
		c.TopPixel()
	}
}
