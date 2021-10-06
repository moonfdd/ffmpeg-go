package ffconstant

/**
 * @addtogroup lavu_video
 * @{
 *
 * @defgroup lavu_video_stereo3d Stereo3D types and functions
 * @{
 */

/**
 * @addtogroup lavu_video_stereo3d
 * A stereoscopic video file consists in multiple views embedded in a single
 * frame, usually describing two views of a scene. This file describes all
 * possible codec-independent view arrangements.
 * */

/**
 * List of possible 3D Types
 */
type AVStereo3DType int32

const (
	/**
	 * Video is not stereoscopic (and metadata has to be there).
	 */
	AV_STEREO3D_2D = 0

	/**
	 * Views are next to each other.
	 *
	 * @code{.unparsed}
	 *    LLLLRRRR
	 *    LLLLRRRR
	 *    LLLLRRRR
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_SIDEBYSIDE

	/**
	 * Views are on top of each other.
	 *
	 * @code{.unparsed}
	 *    LLLLLLLL
	 *    LLLLLLLL
	 *    RRRRRRRR
	 *    RRRRRRRR
	 * @endcode
	 */
	AV_STEREO3D_TOPBOTTOM

	/**
	 * Views are alternated temporally.
	 *
	 * @code{.unparsed}
	 *     frame0   frame1   frame2   ...
	 *    LLLLLLLL RRRRRRRR LLLLLLLL
	 *    LLLLLLLL RRRRRRRR LLLLLLLL
	 *    LLLLLLLL RRRRRRRR LLLLLLLL
	 *    ...      ...      ...
	 * @endcode
	 */
	AV_STEREO3D_FRAMESEQUENCE

	/**
	 * Views are packed in a checkerboard-like structure per pixel.
	 *
	 * @code{.unparsed}
	 *    LRLRLRLR
	 *    RLRLRLRL
	 *    LRLRLRLR
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_CHECKERBOARD

	/**
	 * Views are next to each other, but when upscaling
	 * apply a checkerboard pattern.
	 *
	 * @code{.unparsed}
	 *     LLLLRRRR          L L L L    R R R R
	 *     LLLLRRRR    =>     L L L L  R R R R
	 *     LLLLRRRR          L L L L    R R R R
	 *     LLLLRRRR           L L L L  R R R R
	 * @endcode
	 */
	AV_STEREO3D_SIDEBYSIDE_QUINCUNX

	/**
	 * Views are packed per line, as if interlaced.
	 *
	 * @code{.unparsed}
	 *    LLLLLLLL
	 *    RRRRRRRR
	 *    LLLLLLLL
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_LINES

	/**
	 * Views are packed per column.
	 *
	 * @code{.unparsed}
	 *    LRLRLRLR
	 *    LRLRLRLR
	 *    LRLRLRLR
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_COLUMNS
)

/**
 * List of possible view types.
 */
type AVStereo3DView int32

const (
	/**
	 * Frame contains two packed views.
	 */
	AV_STEREO3D_VIEW_PACKED = 0

	/**
	 * Frame contains only the left view.
	 */
	AV_STEREO3D_VIEW_LEFT

	/**
	 * Frame contains only the right view.
	 */
	AV_STEREO3D_VIEW_RIGHT
)

/**
 * Inverted views, Right/Bottom represents the left view.
 */
const AV_STEREO3D_FLAG_INVERT = (1 << 0)
