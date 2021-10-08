package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"fmt"
	"syscall"
	"unsafe"
)

/**
 * This structure describes how to handle spherical videos, outlining
 * information about projection, initial layout, and any other view modifier.
 *
 * @note The struct must be allocated with av_spherical_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVSphericalMapping struct {
	/**
	 * Projection type.
	 */
	Projection ffconstant.AVSphericalProjection

	/**
	 * @name Initial orientation
	 * @{
	 * There fields describe additional rotations applied to the sphere after
	 * the video frame is mapped onto it. The sphere is rotated around the
	 * viewer, who remains stationary. The order of transformation is always
	 * yaw, followed by pitch, and finally by roll.
	 *
	 * The coordinate system matches the one defined in OpenGL, where the
	 * forward vector (z) is coming out of screen, and it is equivalent to
	 * a rotation matrix of R = r_y(yaw) * r_x(pitch) * r_z(roll).
	 *
	 * A positive yaw rotates the portion of the sphere in front of the viewer
	 * toward their right. A positive pitch rotates the portion of the sphere
	 * in front of the viewer upwards. A positive roll tilts the portion of
	 * the sphere in front of the viewer to the viewer's right.
	 *
	 * These values are exported as 16.16 fixed point.
	 *
	 * See this equirectangular projection as example:
	 *
	 * @code{.unparsed}
	 *                   Yaw
	 *     -180           0           180
	 *   90 +-------------+-------------+  180
	 *      |             |             |                  up
	 * P    |             |             |                 y|    forward
	 * i    |             ^             |                  |   /z
	 * t  0 +-------------X-------------+    0 Roll        |  /
	 * c    |             |             |                  | /
	 * h    |             |             |                 0|/_____right
	 *      |             |             |                        x
	 *  -90 +-------------+-------------+ -180
	 *
	 * X - the default camera center
	 * ^ - the default up vector
	 * @endcode
	 */
	Yaw   ffcommon.FInt32T ///< Rotation around the up vector [-180, 180].
	Pitch ffcommon.FInt32T ///< Rotation around the right vector [-90, 90].
	Roll  ffcommon.FInt32T ///< Rotation around the forward vector [-180, 180].
	/**
	 * @}
	 */

	/**
	 * @name Bounding rectangle
	 * @anchor bounding
	 * @{
	 * These fields indicate the location of the current tile, and where
	 * it should be mapped relative to the original surface. They are
	 * exported as 0.32 fixed point, and can be converted to classic
	 * pixel values with av_spherical_bounds().
	 *
	 * @code{.unparsed}
	 *      +----------------+----------+
	 *      |                |bound_top |
	 *      |            +--------+     |
	 *      | bound_left |tile    |     |
	 *      +<---------->|        |<--->+bound_right
	 *      |            +--------+     |
	 *      |                |          |
	 *      |    bound_bottom|          |
	 *      +----------------+----------+
	 * @endcode
	 *
	 * If needed, the original video surface dimensions can be derived
	 * by adding the current stream or frame size to the related bounds,
	 * like in the following example:
	 *
	 * @code{c}
	 *     original_width  = tile->width  + bound_left + bound_right;
	 *     original_height = tile->height + bound_top  + bound_bottom;
	 * @endcode
	 *
	 * @note These values are valid only for the tiled equirectangular
	 *       projection type (@ref AV_SPHERICAL_EQUIRECTANGULAR_TILE),
	 *       and should be ignored in all other cases.
	 */
	BoundLeft   ffcommon.FUint32T ///< Distance from the left edge
	BoundTop    ffcommon.FUint32T ///< Distance from the top edge
	BoundRight  ffcommon.FUint32T ///< Distance from the right edge
	BoundBottom ffcommon.FUint32T ///< Distance from the bottom edge
	/**
	 * @}
	 */

	/**
	 * Number of pixels to pad from the edge of each cube face.
	 *
	 * @note This value is valid for only for the cubemap projection type
	 *       (@ref AV_SPHERICAL_CUBEMAP), and should be ignored in all other
	 *       cases.
	 */
	Padding ffcommon.FUint32T
}

/**
* Allocate a AVSphericalVideo structure and initialize its fields to default
* values.
*
* @return the newly allocated struct or NULL on failure
 */
//AVSphericalMapping *av_spherical_alloc(size_t *size);
//未测试
func AvSphericalAlloc(size *ffcommon.FSizeT) (res *AVSphericalMapping, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_spherical_alloc").Call(
		uintptr(unsafe.Pointer(size)),
	)
	if err != nil {
		//return
	}
	res = (*AVSphericalMapping)(unsafe.Pointer(t))
	return
}

//
///**
// * Convert the @ref bounding fields from an AVSphericalVideo
// * from 0.32 fixed point to pixels.
// *
// * @param map    The AVSphericalVideo map to read bound values from.
// * @param width  Width of the current frame or stream.
// * @param height Height of the current frame or stream.
// * @param left   Pixels from the left edge.
// * @param top    Pixels from the top edge.
// * @param right  Pixels from the right edge.
// * @param bottom Pixels from the bottom edge.
// */
//void av_spherical_tile_bounds(const AVSphericalMapping *map,
//size_t width, size_t height,
//size_t *left, size_t *top,
//size_t *right, size_t *bottom);
//未测试
func (map0 *AVSphericalMapping) AvSphericalTileBounds(
	width ffcommon.FSizeT, height ffcommon.FSizeT,
	left *ffcommon.FSizeT, top *ffcommon.FSizeT,
	right *ffcommon.FSizeT, bottom *ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_spherical_tile_bounds").Call(
		uintptr(unsafe.Pointer(map0)),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(top)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(bottom)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Provide a human-readable name of a given AVSphericalProjection.
*
* @param projection The input AVSphericalProjection.
*
* @return The name of the AVSphericalProjection, or "unknown".
 */
//const char *av_spherical_projection_name(enum AVSphericalProjection projection);
//未测试
func AvSphericalProjectionName(projection ffconstant.AVSphericalProjection) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_spherical_projection_name").Call(
		uintptr(projection),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

//
///**
// * Get the AVSphericalProjection form a human-readable name.
// *
// * @param name The input string.
// *
// * @return The AVSphericalProjection value, or -1 if not found.
// */
//int av_spherical_from_name(const char *name);
//未测试
func AvSphericalFromName(projection ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var projectionp *byte
	projectionp, err = syscall.BytePtrFromString(projection)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_spherical_projection_name").Call(
		uintptr(unsafe.Pointer(projectionp)),
	)
	if err != nil {
		fmt.Println("err = ", err)
		//return
	}
	res = ffcommon.FInt(t)
	return
}
