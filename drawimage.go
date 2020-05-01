// 30 april 2020

package ui

import (
	"image"
	"unsafe"
)

// #cgo darwin CFLAGS: -fno-objc-arc -mmacosx-version-min=10.8
// #cgo darwin LDFLAGS: -framework Foundation -framework AppKit -mmacosx-version-min=10.8
// const char DrawImage_native(void**, const char*, unsigned int, unsigned int, unsigned int, unsigned int, unsigned int);
import "C"

func (ctx *DrawContext) DrawImage(img *image.Image, x, y uint) {
	rgbaImage := imageToRGBA(*img)
	startPixel := uint(rgbaImage.PixOffset(rgbaImage.Rect.Min.X, rgbaImage.Rect.Min.Y))

	C.DrawImage_native(
		(*unsafe.Pointer)(unsafe.Pointer(ctx)),
		(*C.char)(unsafe.Pointer(&rgbaImage.Pix[startPixel])),
		C.uint(rgbaImage.Rect.Max.X-rgbaImage.Rect.Min.X),
		C.uint(rgbaImage.Rect.Max.Y-rgbaImage.Rect.Min.Y),
		C.uint(rgbaImage.Stride),
		C.uint(x),
		C.uint(y),
	)
}

func imageToRGBA(sourceImage image.Image) (rgbaImage *image.RGBA) {
	size := sourceImage.Bounds().Size()
	rgbaImage = image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	for x := rgbaImage.Rect.Min.X; x < rgbaImage.Rect.Max.X; x++ {
		for y := rgbaImage.Rect.Min.Y; y < rgbaImage.Rect.Max.Y; y++ {
			rgbaImage.Set(x, y, sourceImage.At(x, y))
		}
	}

	return rgbaImage
}
