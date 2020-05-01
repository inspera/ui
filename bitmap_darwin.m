#import <Foundation/Foundation.h>

typedef struct uiDrawContext {
  CGContextRef c;
  CGFloat height;
} uiDrawContext;

typedef struct CGImage uiBitmap;

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

  CGImageRef bmp = CGImageCreate(
      width, height, kBitsPerComponent, kBitsPerComponent * kNumChannels,
      stride, space, info, data, NULL, NO, kCGRenderingIntentDefault);

  CGDataProviderRelease(data);
  CGColorSpaceRelease(space);

  CGContextRetain(ctx->c);

  return bmp;
}

void uiFreeBitmap(uiDrawContext *ctx, uiBitmap *bmp) {
  CGImageRelease(bmp);
  CGContextRelease(ctx->c);
}

void uiDrawBitmap(uiDrawContext *ctx, uiBitmap *bmp, double x, double y) {
  size_t width = CGImageGetWidth(bmp);
  size_t height = CGImageGetHeight(bmp);

  CGContextSaveGState(ctx->c);
  CGContextTranslateCTM(ctx->c, 0, height + 2 * y);
  CGContextScaleCTM(ctx->c, 1, -1);
  CGContextDrawImage(ctx->c, CGRectMake(x, y, width, height), bmp);
  CGContextRestoreGState(ctx->c);
}
