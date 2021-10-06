package ffconstant

/**
 * Pixel format is big-endian.
 */
const AV_PIX_FMT_FLAG_BE = (1 << 0)

/**
 * Pixel format has a palette in data[1], values are indexes in this palette.
 */
const AV_PIX_FMT_FLAG_PAL = (1 << 1)

/**
 * All values of a component are bit-wise packed end to end.
 */
const AV_PIX_FMT_FLAG_BITSTREAM = (1 << 2)

/**
 * Pixel format is an HW accelerated format.
 */
const AV_PIX_FMT_FLAG_HWACCEL = (1 << 3)

/**
 * At least one pixel component is not in the first data plane.
 */
const AV_PIX_FMT_FLAG_PLANAR = (1 << 4)

/**
 * The pixel format contains RGB-like data (as opposed to YUV/grayscale).
 */
const AV_PIX_FMT_FLAG_RGB = (1 << 5)

//#if FF_API_PSEUDOPAL
///**
// * The pixel format is "pseudo-paletted". This means that it contains a
// * fixed palette in the 2nd plane but the palette is fixed/ffconstant for each
// * PIX_FMT. This allows interpreting the data as if it was PAL8, which can
// * in some cases be simpler. Or the data can be interpreted purely based on
// * the pixel format without using the palette.
// * An example of a pseudo-paletted format is AV_PIX_FMT_GRAY8
// *
// * @deprecated This flag is deprecated, and will be removed. When it is removed,
// * the extra palette allocation in AVFrame.data[1] is removed as well. Only
// * actual paletted formats (as indicated by AV_PIX_FMT_FLAG_PAL) will have a
// * palette. Starting with FFmpeg versions which have this flag deprecated, the
// * extra "pseudo" palette is already ignored, and API users are not required to
// * allocate a palette for AV_PIX_FMT_FLAG_PSEUDOPAL formats (it was required
// * before the deprecation, though).
// */
//#define AV_PIX_FMT_FLAG_PSEUDOPAL    (1 << 6)
//#endif

/**
 * The pixel format has an alpha channel. This is set on all formats that
 * support alpha in some way, including AV_PIX_FMT_PAL8. The alpha is always
 * straight, never pre-multiplied.
 *
 * If a codec or a filter does not support alpha, it should set all alpha to
 * opaque, or use the equivalent pixel formats without alpha component, e.g.
 * AV_PIX_FMT_RGB0 (or AV_PIX_FMT_RGB24 etc.) instead of AV_PIX_FMT_RGBA.
 */
const AV_PIX_FMT_FLAG_ALPHA = (1 << 7)

/**
 * The pixel format is following a Bayer pattern
 */
const AV_PIX_FMT_FLAG_BAYER = (1 << 8)

/**
 * The pixel format contains IEEE-754 floating point values. Precision (double,
 * single, or half) should be determined by the pixel size (64, 32, or 16 bits).
 */
const AV_PIX_FMT_FLAG_FLOAT = (1 << 9)

const FF_LOSS_RESOLUTION = 0x0001 /**< loss due to resolution change */
const FF_LOSS_DEPTH = 0x0002      /**< loss due to color depth change */
const FF_LOSS_COLORSPACE = 0x0004 /**< loss due to color space conversion */
const FF_LOSS_ALPHA = 0x0008      /**< loss of alpha bits */
const FF_LOSS_COLORQUANT = 0x0010 /**< loss due to color quantization */
const FF_LOSS_CHROMA = 0x0020     /**< loss of chroma (e.g. RGB to gray conversion) */
