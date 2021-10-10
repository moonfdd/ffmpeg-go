package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @addtogroup lavu_video
 * @{
 *
 * @defgroup lavu_video_display Display transformation matrix functions
 * @{
 */

/**
 * @addtogroup lavu_video_display
 * The display transformation matrix specifies an affine transformation that
 * should be applied to video frames for correct presentation. It is compatible
 * with the matrices stored in the ISO/IEC 14496-12 container format.
 *
 * The data is a 3x3 matrix represented as a 9-element array:
 *
 * @code{.unparsed}
 *                                  | a b u |
 *   (a, b, u, c, d, v, x, y, w) -> | c d v |
 *                                  | x y w |
 * @endcode
 *
 * All numbers are stored in native endianness, as 16.16 fixed-point values,
 * except for u, v and w, which are stored as 2.30 fixed-point values.
 *
 * The transformation maps a point (p, q) in the source (pre-transformation)
 * frame to the point (p', q') in the destination (post-transformation) frame as
 * follows:
 *
 * @code{.unparsed}
 *               | a b u |
 *   (p, q, 1) . | c d v | = z * (p', q', 1)
 *               | x y w |
 * @endcode
 *
 * The transformation can also be more explicitly written in components as
 * follows:
 *
 * @code{.unparsed}
 *   p' = (a * p + c * q + x) / z;
 *   q' = (b * p + d * q + y) / z;
 *   z  =  u * p + v * q + w
 * @endcode
 */

/**
 * Extract the rotation component of the transformation matrix.
 *
 * @param matrix the transformation matrix
 * @return the angle (in degrees) by which the transformation rotates the frame
 *         counterclockwise. The angle will be in range [-180.0, 180.0],
 *         or NaN if the matrix is singular.
 *
 * @note floating point numbers are inherently inexact, so callers are
 *       recommended to round the return value to nearest integer before use.
 */
//double av_display_rotation_get(const int32_t matrix[9]);
//未测试
func AvDisplayRotationGet(matrix [9]ffcommon.FInt32T) (res ffcommon.FDouble, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_display_rotation_get").Call(
		uintptr(unsafe.Pointer(&matrix)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FDouble(t)
	return
}

/**
 * Initialize a transformation matrix describing a pure counterclockwise
 * rotation by the specified angle (in degrees).
 *
 * @param matrix an allocated transformation matrix (will be fully overwritten
 *               by this function)
 * @param angle rotation angle in degrees.
 */
//void av_display_rotation_set(int32_t matrix[9], double angle);
//未测试
func AvDisplayRotationSet(matrix [9]ffcommon.FInt32T, angle ffcommon.FDouble) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_display_rotation_set").Call(
		uintptr(unsafe.Pointer(&matrix)),
		uintptr(angle),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Flip the input matrix horizontally and/or vertically.
 *
 * @param matrix an allocated transformation matrix
 * @param hflip whether the matrix should be flipped horizontally
 * @param vflip whether the matrix should be flipped vertically
 */
//void av_display_matrix_flip(int32_t matrix[9], int hflip, int vflip);
//未测试
func AvDisplayMatrixFlip(matrix [9]ffcommon.FInt32T, hflip ffcommon.FInt, vflip ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_display_matrix_flip").Call(
		uintptr(unsafe.Pointer(&matrix)),
		uintptr(hflip),
		uintptr(vflip),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
