#ifdef WIN32
#include <iostream>

#include <d2d1.h>
#include <d2d1helper.h>

// void(document.getElementById("editor").env.editor.session.setMode("ace/mode/c_cpp"))

extern "C" const char insperaDrawImage(void**,  const char*, unsigned int, unsigned int, unsigned int, unsigned int, unsigned int);

struct insperaUiDrawContext {
    ID2D1RenderTarget *rt;
};

const char insperaDrawImage(
    void** theContext,
    const char bunnyImage[],
    unsigned int width,
    unsigned int height,
    unsigned int stride,
    unsigned int x,
    unsigned int y
) {
    struct insperaUiDrawContext *context = NULL;
    ID2D1Bitmap *bitmap = NULL;
    D2D1_PIXEL_FORMAT pixelFormat;
    context = (struct insperaUiDrawContext*)(*theContext);

    pixelFormat = D2D1::PixelFormat(
        DXGI_FORMAT_B8G8R8A8_UNORM,
        D2D1_ALPHA_MODE_PREMULTIPLIED
    );

    context->rt->CreateBitmap(
        D2D1::SizeU(width, height),
        bunnyImage,
        stride,
        {pixelFormat, 96.0f, 96.0f},
        &bitmap
    );

    context->rt->DrawBitmap(
        bitmap,
        D2D1::RectF(
            (float)(x),
            (float)(y),
            (float)(width + x),
            (float)(height + y)
        ),
        1.0f,
        D2D1_BITMAP_INTERPOLATION_MODE_LINEAR,
        NULL
    );

    // Once drawn, release it.
    // TODO: Create a bitmap object and let Go handle its lifecycle.
    bitmap->Release();

    return bunnyImage[0];
}

#endif
