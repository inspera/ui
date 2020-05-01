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
// void uiFreeBitmap(uiDrawContext *ctx, uiBitmap *bmp);
// void uiDrawBitmap(uiDrawContext *ctx, uiBitmap *bmp, double x, double y);
//
import "C"

import (
	"image"
	"image/draw"
	"unsafe"
)

// Bitmap represents a bitmap capable to be drawn on drawing context.
type Bitmap struct {
	ctx *DrawContext
	bmp *C.uiBitmap
}

// NewBitmap creates a new bitmap from a given image. The resulting bitmap is
// associated with the current drawing context.
func (c *DrawContext) NewBitmap(img image.Image) *Bitmap {
	bounds := img.Bounds()
	rgba, ok := img.(*image.RGBA)
	if !ok {
		rgba = image.NewRGBA(bounds)
		draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	}

	bmp := C.uiNewBitmap(c.c, C.int(bounds.Dx()), C.int(bounds.Dy()),
		C.int(rgba.Stride), unsafe.Pointer(&rgba.Pix[0]))
	if bmp == nil {
		panic("failed to create a bitmap")
	}

	return &Bitmap{c, bmp}
}

// Free destroys the given bitmap.
func (b *Bitmap) Free() {
	C.uiFreeBitmap(b.ctx.c, b.bmp)
}

// Draw draws the bitmap on its drawing context.
func (b *Bitmap) Draw(x, y float64) {
	C.uiDrawBitmap(b.ctx.c, b.bmp, C.double(x), C.double(y))
}

// DrawImage is a convenience shortcut to create and draw a disposable bitmap.
func (c *DrawContext) DrawImage(img image.Image, x, y float64) {
	bmp := c.NewBitmap(img)
	defer bmp.Free()
	bmp.Draw(x, y)
}
