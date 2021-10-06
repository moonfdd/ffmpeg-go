package ffconstant

/**
 * @addtogroup version_utils
 *
 * Useful to check and match library version in order to maintain
 * backward compatibility.
 *
 * The FFmpeg libraries follow a versioning sheme very similar to
 * Semantic Versioning (http://semver.org/)
 * The difference is that the component called PATCH is called MICRO in FFmpeg
 * and its value is reset to 100 instead of 0 to keep it above or equal to 100.
 * Also we do not increase MICRO for every bugfix or change in git master.
 *
 * Prior to FFmpeg 3.2 point releases did not change any lib version number to
 * avoid aliassing different git master checkouts.
 * Starting with FFmpeg 3.2, the released library versions will occupy
 * a separate MAJOR.MINOR that is not used on the master development branch.
 * That is if we branch a release of master 55.10.123 we will bump to 55.11.100
 * for the release and master will continue at 55.12.100 after it. Each new
 * point release will then bump the MICRO improving the usefulness of the lib
 * versions.
 *
 * @{
 */

//const AV_VERSION_INT(a, b, c) ((a)<<16 | (b)<<8 | (c))
//const AV_VERSION_DOT(a, b, c) a ##.## b ##.## c
//const AV_VERSION(a, b, c) AV_VERSION_DOT(a, b, c)

/**
 * Extract version components from the full ::AV_VERSION_INT int as returned
 * by functions like ::avformat_version() and ::avcodec_version()
 */
//const AV_VERSION_MAJOR(a) ((a) >> 16)
//const AV_VERSION_MINOR(a) (((a) & 0x00FF00) >> 8)
//const AV_VERSION_MICRO(a) ((a) & 0xFF)

/**
 * @}
 */

/**
 * @defgroup lavu_ver Version and Build diagnostics
 *
 * Macros and function useful to check at compiletime and at runtime
 * which version of libavutil is in use.
 *
 * @{
 */

const LIBAVUTIL_VERSION_MAJOR = 56
const LIBAVUTIL_VERSION_MINOR = 70
const LIBAVUTIL_VERSION_MICRO = 100

//const LIBAVUTIL_VERSION_INT = 1 //AV_VERSION_INT(LIBAVUTIL_VERSION_MAJOR, \
//LIBAVUTIL_VERSION_MINOR, \
//LIBAVUTIL_VERSION_MICRO)
//const LIBAVUTIL_VERSION = 1 //  AV_VERSION(LIBAVUTIL_VERSION_MAJOR,     \
//LIBAVUTIL_VERSION_MINOR,     \
//LIBAVUTIL_VERSION_MICRO)
//const LIBAVUTIL_BUILD = LIBAVUTIL_VERSION_INT

const LIBAVUTIL_IDENT = "" //  "Lavu" AV_STRINGIFY(LIBAVUTIL_VERSION)

/**
 * @defgroup lavu_depr_guards Deprecation Guards
 * FF_API_* defines may be placed below to indicate public API that will be
 * dropped at a future version bump. The defines themselves are not part of
 * the public API and may change, break or disappear at any time.
 *
 * @note, when bumping the major version it is recommended to manually
 * disable each FF_API_* in its own commit instead of disabling them all
 * at once through the bump. This improves the git bisect-ability of the change.
 *
 * @{
 */

const FF_API_VAAPI = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_FRAME_QP = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_PLUS1_MINUS1 = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_ERROR_FRAME = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_PKT_PTS = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_CRYPTO_SIZE_T = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_FRAME_GET_SET = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_PSEUDOPAL = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_CHILD_CLASS_NEXT = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_BUFFER_SIZE_T = (LIBAVUTIL_VERSION_MAJOR < 57)
const FF_API_D2STR = (LIBAVUTIL_VERSION_MAJOR < 58)
const FF_API_DECLARE_ALIGNED = (LIBAVUTIL_VERSION_MAJOR < 58)
