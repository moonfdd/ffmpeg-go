package ffconstant

type AVEscapeMode int32

const (
	AV_ESCAPE_MODE_AUTO      = 0 ///< Use auto-selected escaping mode.
	AV_ESCAPE_MODE_BACKSLASH     ///< Use backslash escaping.
	AV_ESCAPE_MODE_QUOTE         ///< Use single-quote escaping.
	AV_ESCAPE_MODE_XML           ///< Use XML non-markup character data escaping.
)

/**
 * Consider spaces special and escape them even in the middle of the
 * string.
 *
 * This is equivalent to adding the whitespace characters to the special
 * characters lists, except it is guaranteed to use the exact same list
 * of whitespace characters as the rest of libavutil.
 */
const AV_ESCAPE_FLAG_WHITESPACE = (1 << 0)

/**
 * Escape only specified special characters.
 * Without this flag, escape also any characters that may be considered
 * special by av_get_token(), such as the single quote.
 */
const AV_ESCAPE_FLAG_STRICT = (1 << 1)

/**
 * Within AV_ESCAPE_MODE_XML, additionally escape single quotes for single
 * quoted attributes.
 */
const AV_ESCAPE_FLAG_XML_SINGLE_QUOTES = (1 << 2)

/**
 * Within AV_ESCAPE_MODE_XML, additionally escape double quotes for double
 * quoted attributes.
 */
const AV_ESCAPE_FLAG_XML_DOUBLE_QUOTES = (1 << 3)

const AV_UTF8_FLAG_ACCEPT_INVALID_BIG_CODES = 1          ///< accept codepoints over 0x10FFFF
const AV_UTF8_FLAG_ACCEPT_NON_CHARACTERS = 2             ///< accept non-characters - 0xFFFE and 0xFFFF
const AV_UTF8_FLAG_ACCEPT_SURROGATES = 4                 ///< accept UTF-16 surrogates codes
const AV_UTF8_FLAG_EXCLUDE_XML_INVALID_CONTROL_CODES = 8 ///< exclude control codes not accepted by XML

const AV_UTF8_FLAG_ACCEPT_ALL = AV_UTF8_FLAG_ACCEPT_INVALID_BIG_CODES | AV_UTF8_FLAG_ACCEPT_NON_CHARACTERS | AV_UTF8_FLAG_ACCEPT_SURROGATES
