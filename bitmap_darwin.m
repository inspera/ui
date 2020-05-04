#import <Foundation/Foundation.h>

typedef struct uiDrawContext {
  CGContextRef c;
  CGFloat height;
} uiDrawContext;

typedef struct {
  CGContextRef ctx;
  CGImageRef img;
} uiBitmap;

uiBitmap *uiNewBitmap(uiDrawContext *ctx, int width, int height, int stride,
                      const void *rgba) {
  CGColorSpaceRef space = CGColorSpaceCreateDeviceRGB();
  if (space == NULL) {
    return NULL;
  }

  CGDataProviderRef data =
      CGDataProviderCreateWithData(NULL, rgba, stride * height, NULL);

  const int kBitsPerComponent = 8;
  const int kNumChannels = 4;

  CGBitmapInfo info =
      kCGImageAlphaPremultipliedFirst | kCGBitmapByteOrder32Little;

  CGImageRef img = CGImageCreate(
      width, height, kBitsPerComponent, kBitsPerComponent * kNumChannels,
      stride, space, info, data, NULL, NO, kCGRenderingIntentDefault);

  CGDataProviderRelease(data);
  CGColorSpaceRelease(space);

  CGContextRetain(ctx->c);

  uiBitmap *bmp = malloc(sizeof(uiBitmap));
  bmp->ctx = ctx->c;
  bmp->img = img;

  return bmp;
}

void uiFreeBitmap(uiBitmap *bmp) {
  CGImageRelease(bmp->img);
  CGContextRelease(bmp->ctx);
  free(bmp);
}

void uiDrawBitmap(uiBitmap *bmp, double x, double y) {
  size_t width = CGImageGetWidth(bmp->img);
  size_t height = CGImageGetHeight(bmp->img);

  CGContextSaveGState(bmp->ctx);
  CGContextTranslateCTM(bmp->ctx, 0, height + 2 * y);
  CGContextScaleCTM(bmp->ctx, 1, -1);
  CGContextDrawImage(bmp->ctx, CGRectMake(x, y, width, height), bmp->img);
  CGContextRestoreGState(bmp->ctx);
}
