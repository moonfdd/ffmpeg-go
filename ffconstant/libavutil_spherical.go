package ffconstant

/**
 * @addtogroup lavu_video
 * @{
 *
 * @defgroup lavu_video_spherical Spherical video mapping
 * @{
 */

/**
 * @addtogroup lavu_video_spherical
 * A spherical video file contains surfaces that need to be mapped onto a
 * sphere. Depending on how the frame was converted, a different distortion
 * transformation or surface recomposition function needs to be applied before
 * the video should be mapped and displayed.
 */

/**
 * Projection of the video surface(s) on a sphere.
 */
type AVSphericalProjection int32

const (
	/**
	 * Video represents a sphere mapped on a flat surface using
	 * equirectangular projection.
	 */
	AV_SPHERICAL_EQUIRECTANGULAR = 0

	/**
	 * Video frame is split into 6 faces of a cube, and arranged on a
	 * 3x2 layout. Faces are oriented upwards for the front, left, right,
	 * and back faces. The up face is oriented so the top of the face is
	 * forwards and the down face is oriented so the top of the face is
	 * to the back.
	 */
	AV_SPHERICAL_CUBEMAP

	/**
	 * Video represents a portion of a sphere mapped on a flat surface
	 * using equirectangular projection. The @ref bounding fields indicate
	 * the position of the current video in a larger surface.
	 */
	AV_SPHERICAL_EQUIRECTANGULAR_TILE
)
