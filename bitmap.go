// +build windows,cgo darwin,cgo
package ui

// #cgo darwin CFLAGS: -mmacosx-version-min=10.12
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.12
//
// typedef struct uiDrawContext uiDrawContext;
// typedef struct uiBitmap uiBitmap;
//
// uiBitmap *uiNewBitmap(uiDrawContext *ctx, int width, int height, int stride,
//                       const void *rgba);
// void uiFreeBitmap(uiBitmap *bmp);
// void uiDrawBitmap(uiBitmap *bmp, double x, double y);
//
// int uiDrawImage(uiDrawContext *ctx, int width, int height, int stride,
//                 const void *rgba, double x, double y) {
//   uiBitmap *bmp = uiNewBitmap(ctx, width, height, stride, rgba);
//   if (bmp) {
//     uiDrawBitmap(bmp, x, y);
//     uiFreeBitmap(bmp);
//   }
//   return !!bmp;
// }
//
import "C"

import (
	"image"
	"image/draw"
	"unsafe"
)

// Bitmap represents a bitmap capable to be drawn on drawing context.
type Bitmap struct {
	b *C.uiBitmap
}

func imageToRGBAData(img image.Image) (C.int, C.int, C.int, unsafe.Pointer) {
	bounds := img.Bounds()
	rgba, ok := img.(*image.RGBA)
	if !ok {
		rgba = image.NewRGBA(bounds)
		draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	}
	return C.int(bounds.Dx()), C.int(bounds.Dy()),
		C.int(rgba.Stride), unsafe.Pointer(&rgba.Pix[0])
}

// NewBitmap creates a new bitmap from a given image. The resulting bitmap is
// associated with the current drawing context.
func (c *DrawContext) NewBitmap(img image.Image) *Bitmap {
	width, height, stride, rgba := imageToRGBAData(img)

	bmp := C.uiNewBitmap(c.c, width, height, stride, rgba)
	if bmp == nil {
		panic("failed to create a bitmap")
	}

	return &Bitmap{bmp}
}

// Free destroys the given bitmap.
func (b *Bitmap) Free() {
	C.uiFreeBitmap(b.b)
}

// Draw draws the bitmap on its drawing context.
func (b *Bitmap) Draw(x, y float64) {
	C.uiDrawBitmap(b.b, C.double(x), C.double(y))
}

// DrawImage is a shortcut to create and draw a disposable bitmap.
func (c *DrawContext) DrawImage(img image.Image, x, y float64) {
	width, height, stride, rgba := imageToRGBAData(img)
	if C.uiDrawImage(c.c, width, height, stride, rgba,
		C.double(x), C.double(y)) == 0 {
		panic("failed to draw an image")
	}
}
