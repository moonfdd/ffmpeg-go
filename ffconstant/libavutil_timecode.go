package ffconstant

const AV_TIMECODE_STR_SIZE = 23

type AVTimecodeFlag int32

const (
	AV_TIMECODE_FLAG_DROPFRAME     = 1 << 0 ///< timecode is drop frame
	AV_TIMECODE_FLAG_24HOURSMAX    = 1 << 1 ///< timecode wraps after 24 hours
	AV_TIMECODE_FLAG_ALLOWNEGATIVE = 1 << 2 ///< negative time values are allowed
)
