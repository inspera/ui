// +build !windows,!darwin

package ui

import (
	"image"
)

func DrawImage(context *DrawContext, rgbaImage *image.RGBA, x, y uint) {
	return
}
