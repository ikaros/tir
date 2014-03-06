package render

import (
	"image"
	"image/color"
)

type Renderer interface {
	Render(img image.Image) string
	RenderBig(img image.Image) string
}

func IsTransparent(cl color.Color) bool {
	_, _, _, a := cl.RGBA()
	return a == 0
}
