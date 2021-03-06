#include <cstdio>
#include <vector>

// See
// https://stackoverflow.com/questions/27888109/rendertarget-getsize-not-working.
#define WIDL_EXPLICIT_AGGREGATE_RETURNS
#include <d2d1.h>

#include <comdef.h>

struct uiDrawContext {
  ID2D1RenderTarget *rt;
  std::vector<struct drawState> *states;
  ID2D1PathGeometry *currentClip;
};

typedef struct uiBitmap {
  ID2D1RenderTarget *rt;
  ID2D1Bitmap *img;
} uiBitmap;

extern "C" {

uiBitmap *uiNewBitmap(uiDrawContext *ctx, int width, int height, int stride,
                      const void *rgba) {
  auto fmt = D2D1::PixelFormat(DXGI_FORMAT_B8G8R8A8_UNORM,
                               D2D1_ALPHA_MODE_PREMULTIPLIED);

  float dpi_x, dpi_y;
  ctx->rt->GetDpi(&dpi_x, &dpi_y);

  ID2D1Bitmap *img;
  auto res = ctx->rt->CreateBitmap(D2D1::SizeU(width, height), rgba, stride,
                                   {fmt, dpi_x, dpi_y}, &img);

  if (res != S_OK) {
    // TODO: Pass an error message to the caller.
    fprintf(stderr, "failed to create bitmap: %s\n",
            _com_error(res).ErrorMessage());
    return nullptr;
  }

  return new uiBitmap{ctx->rt, img};
}

void uiFreeBitmap(uiBitmap *bmp) {
  bmp->img->Release();
  delete bmp;
}

void uiDrawBitmap(uiBitmap *bmp, double x, double y, double width,
                  double height) {
  auto size = bmp->img->GetSize();
  width = width > 0 ? width : size.width;
  height = height > 0 ? height : size.height;

  D2D1_RECT_F rect{static_cast<float>(x), static_cast<float>(y),
                   static_cast<float>(width + x),
                   static_cast<float>(height + y)};
  bmp->rt->DrawBitmap(bmp->img, rect, 1,
                      D2D1_BITMAP_INTERPOLATION_MODE_NEAREST_NEIGHBOR, nullptr);
}

}
