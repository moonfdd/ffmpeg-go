package ffconstant

/**
 * Maximum value that av_hash_get_size() will currently return.
 *
 * You can use this if you absolutely want or need to use static allocation for
 * the output buffer and are fine with not supporting hashes newly added to
 * libavutil without recompilation.
 *
 * @warning
 * Adding new hashes with larger sizes, and increasing the macro while doing
 * so, will not be considered an ABI change. To prevent your code from
 * overflowing a buffer, either dynamically allocate the output buffer with
 * av_hash_get_size(), or limit your use of the Hashing API to hashes that are
 * already in FFmpeg during the time of compilation.
 */
const AV_HASH_MAX_SIZE = 64
