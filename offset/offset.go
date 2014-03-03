package offset

import (
	"bytes"
)

var (
	X string
	Y string
)

func Init(oX int, oY int) {
	X = genOffset(oX, ' ')
	Y = genOffset(oY, '\n')
}

func genOffset(n int, c byte) string {
	var b bytes.Buffer

	for i := 0; i < n; i++ {
		b.WriteByte(c)
	}

	return b.String()
}
