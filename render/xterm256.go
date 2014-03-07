package render

import (
	"bytes"
	"image"
	"image/color"
	"strconv"

	"github.com/ikaros/tir/colors"
	"github.com/ikaros/tir/offset"
)

type XTerm256 struct {
	B bytes.Buffer
}

func (r *XTerm256) Render(img image.Image) string {
	bounds := img.Bounds()

	r.B.WriteString(offset.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {

		r.B.WriteString(offset.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r.Pixel(img.At(x, y), img.At(x, y+1))
		}

		r.NewLine()
	}

	return r.B.String()
}

func (r *XTerm256) RenderBig(img image.Image) string {
	bounds := img.Bounds()

	r.B.WriteString(offset.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

		r.B.WriteString(offset.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r.BigPixel(img.At(x, y))
		}

		r.NewLine()
	}

	return r.B.String()
}

func (r *XTerm256) FG(colorCode int) {
	r.B.WriteString("\033[38;5;")
	r.B.WriteString(strconv.Itoa(colorCode))
	r.B.WriteString("m")
}

func (r *XTerm256) BG(colorCode int) {
	r.B.WriteString("\033[48;5;")
	r.B.WriteString(strconv.Itoa(colorCode))
	r.B.WriteString("m")
}

func (r *XTerm256) TopPixel() {
	r.B.WriteString("▀")
}

func (r *XTerm256) BottomPixel() {
	r.B.WriteString("▄")
}

func (r *XTerm256) EmptyPixel() {
	r.B.WriteString(" ")
}

func (r *XTerm256) Reset() {
	r.B.WriteString("\033[0m")
}

func (r *XTerm256) NewLine() {
	r.Reset()
	r.B.WriteString("\n")
}

func (r *XTerm256) BigPixel(c color.Color) {
	r.Reset()

	if !IsTransparent(c) {
		r.BG(colors.XTerm256.Index(c))
	}

	r.B.WriteString("  ")
}

func (r *XTerm256) Pixel(top color.Color, bottom color.Color) {
	r.Reset()

	if IsTransparent(top) && IsTransparent(bottom) {
		r.EmptyPixel()

	} else if IsTransparent(top) {
		r.FG(colors.XTerm256.Index(bottom))
		r.BottomPixel()

	} else if IsTransparent(bottom) {
		r.FG(colors.XTerm256.Index(top))
		r.TopPixel()

	} else {
		r.FG(colors.XTerm256.Index(top))
		r.BG(colors.XTerm256.Index(bottom))
		r.TopPixel()
	}
}
