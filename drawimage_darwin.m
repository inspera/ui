#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>

/* javascript:void(document.getElementById("editor").env.editor.session.setMode("ace/mode/objectivec")) */

struct insperaUiDrawContext {
    CGContextRef OSContext;
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
    const size_t bitsPerChannel = 8;
    const size_t channelCount   = 4;

    struct insperaUiDrawContext *context = NULL;
    size_t imageSize = 0;
    CGDataProviderRef bitmapDataProvider = NULL;
    CGImageRef bitmap = NULL;
    CGColorSpaceRef rgbColorSpace = NULL;

    context = (struct insperaUiDrawContext*)(*theContext);
    imageSize = width * height * channelCount;

    // Must be released manually
    bitmapDataProvider = CGDataProviderCreateWithData(
        NULL,
        bunnyImage,
        imageSize,
        NULL
    );

    if (bitmapDataProvider == NULL) {
        return -1;
    }

    // Must be released manually
    rgbColorSpace = CGColorSpaceCreateDeviceRGB();

    if (rgbColorSpace == NULL) {
        CGDataProviderRelease(bitmapDataProvider);
        return -1;
    }

    // Must be released manually
    bitmap = CGImageCreate(
        (size_t)width,
        (size_t)height,
        bitsPerChannel,
        bitsPerChannel * channelCount,
        stride,
        rgbColorSpace,
        (kCGImageAlphaPremultipliedFirst |
            kCGBitmapByteOrder32Little),
        bitmapDataProvider,
        NULL,
        TRUE,
        kCGRenderingIntentDefault
    );

    if (bitmap == NULL) {
        CGDataProviderRelease(bitmapDataProvider);
        CGColorSpaceRelease(rgbColorSpace);
        return -1;

    }

    CGContextSaveGState(context->OSContext);
    CGContextTranslateCTM(context->OSContext, (CGFloat)0, (CGFloat)(height + 2*y));
    CGContextScaleCTM(context->OSContext, 1.0, -1.0);

    CGContextDrawImage(
        context->OSContext,
        CGRectMake(
            (CGFloat)x,
            (CGFloat)y,
            (CGFloat)width,
            (CGFloat)height
        ),
        bitmap
    );

    CGContextRestoreGState(context->OSContext);
    CGImageRelease(bitmap);
    CGDataProviderRelease(bitmapDataProvider);
    CGColorSpaceRelease(rgbColorSpace);

    return 0;

}
