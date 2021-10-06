package libavformat

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
)

/**
 * Callback for checking whether to abort blocking functions.
 * AVERROR_EXIT is returned in this case by the interrupted
 * function. During blocking operations, callback is called with
 * opaque as parameter. If the callback returns 1, the
 * blocking operation will be aborted.
 *
 * No members can be added to this struct without a major bump, if
 * new elements have been added after this struct in AVFormatContext
 * or AVIOContext.
 */
type AVIOInterruptCB struct {
}

/**
 * Describes single entry of the directory.
 *
 * Only name and type fields are guaranteed be set.
 * Rest of fields are protocol or/and platform dependent and might be unknown.
 */
type AVIODirEntry struct {
}

type AVIODirContext struct {
}

/**
 * Bytestream IO Context.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVIOContext) must not be used outside libav*.
 *
 * @note None of the function pointers in AVIOContext should be called
 *       directly, they should only be set by the client application
 *       when implementing custom I/O. Normally these are set to the
 *       function pointers specified in avio_alloc_context()
 */
type AVIOContext struct {
}

/**
 * Return the name of the protocol that will handle the passed URL.
 *
 * NULL is returned if no protocol could be found for the given URL.
 *
 * @return Name of the protocol or NULL.
 */
//const char *avio_find_protocol_name(const char *url);
func AvioFindProtocolName(url ffcommon.FConstCharP) (res ffcommon.FConstCharP) {
	return
}

/**
 * Return AVIO_FLAG_* access flags corresponding to the access permissions
 * of the resource in url, or a negative value corresponding to an
 * AVERROR code in case of failure. The returned access flags are
 * masked by the value in flags.
 *
 * @note This function is intrinsically unsafe, in the sense that the
 * checked resource may change its existence or permission status from
 * one call to another. Thus you should not trust the returned value,
 * unless you are sure that no other processes are accessing the
 * checked resource.
 */
//int avio_check(const char *url, int flags);
func AvioCheck(url ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Move or rename a resource.
 *
 * @note url_src and url_dst should share the same protocol and authority.
 *
 * @param url_src url to resource to be moved
 * @param url_dst new url to resource if the operation succeeded
 * @return >=0 on success or negative on error.
 */
//int avpriv_io_move(const char *url_src, const char *url_dst);
func AvprivIoMove(url_src ffcommon.FConstCharP, url_dst ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Delete a resource.
 *
 * @param url resource to be deleted.
 * @return >=0 on success or negative on error.
 */
//int avpriv_io_delete(const char *url);
func AvprivIoDelete(url ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Open directory for reading.
 *
 * @param s       directory read context. Pointer to a NULL pointer must be passed.
 * @param url     directory to be listed.
 * @param options A dictionary filled with protocol-private options. On return
 *                this parameter will be destroyed and replaced with a dictionary
 *                containing options that were not found. May be NULL.
 * @return >=0 on success or negative on error.
 */
//int avio_open_dir(AVIODirContext **s, const char *url, AVDictionary **options);
func AvioOpenDir(s **AVIODirContext, url ffcommon.FConstCharP, options **AVDictionary) (res ffcommon.FInt) {
	return
}

/**
 * Get next directory entry.
 *
 * Returned entry must be freed with avio_free_directory_entry(). In particular
 * it may outlive AVIODirContext.
 *
 * @param s         directory read context.
 * @param[out] next next entry or NULL when no more entries.
 * @return >=0 on success or negative on error. End of list is not considered an
 *             error.
 */
//int avio_read_dir(AVIODirContext *s, AVIODirEntry **next);
func (s *AVIODirContext) AvioReadDir(next **AVDictionary) (res ffcommon.FInt) {
	return
}

/**
 * Close directory.
 *
 * @note Entries created using avio_read_dir() are not deleted and must be
 * freeded with avio_free_directory_entry().
 *
 * @param s         directory read context.
 * @return >=0 on success or negative on error.
 */
//int avio_close_dir(AVIODirContext **s);
func AvioCloseDir(s **AVIODirContext) (res ffcommon.FInt) {
	return
}

/**
 * Free entry allocated by avio_read_dir().
 *
 * @param entry entry to be freed.
 */
//void avio_free_directory_entry(AVIODirEntry **entry);
func AvioFreeDirectoryEntry(entry **AVIODirEntry) {
	return
}

/**
 * Allocate and initialize an AVIOContext for buffered I/O. It must be later
 * freed with avio_context_free().
 *
 * @param buffer Memory block for input/output operations via AVIOContext.
 *        The buffer must be allocated with av_malloc() and friends.
 *        It may be freed and replaced with a new buffer by libavformat.
 *        AVIOContext.buffer holds the buffer currently in use,
 *        which must be later freed with av_free().
 * @param buffer_size The buffer size is very important for performance.
 *        For protocols with fixed blocksize it should be set to this blocksize.
 *        For others a typical size is a cache page, e.g. 4kb.
 * @param write_flag Set to 1 if the buffer should be writable, 0 otherwise.
 * @param opaque An opaque pointer to user-specific data.
 * @param read_packet  A function for refilling the buffer, may be NULL.
 *                     For stream protocols, must never return 0 but rather
 *                     a proper AVERROR code.
 * @param write_packet A function for writing the buffer contents, may be NULL.
 *        The function may not change the input buffers content.
 * @param seek A function for seeking to specified byte position, may be NULL.
 *
 * @return Allocated AVIOContext or NULL on failure.
 */
//AVIOContext *avio_alloc_context(
//unsigned char *buffer,
//int buffer_size,
//int write_flag,
//void *opaque,
//int (*read_packet)(void *opaque, uint8_t *buf, int buf_size),
//int (*write_packet)(void *opaque, uint8_t *buf, int buf_size),
//int64_t (*seek)(void *opaque, int64_t offset, int whence));
func AvioAllocContext(
	buffer ffcommon.FUnsignedCharP,
	buffer_size ffcommon.FInt,
	write_flag ffcommon.FInt,
	opaque ffcommon.FVoidP,
	//int (*read_packet)(void *opaque, uint8_t *buf, int buf_size),
	//int (*write_packet)(void *opaque, uint8_t *buf, int buf_size),
	//int64_t (*seek)(void *opaque, int64_t offset, int whence)
) (res *AVIOContext) {
	//未完成
	return
}

/**
 * Free the supplied IO context and everything associated with it.
 *
 * @param s Double pointer to the IO context. This function will write NULL
 * into s.
 */
//void avio_context_free(AVIOContext **s);
func (s *AVIOContext) AvioContextFree() {
	//s再取地址
	return
}

//void avio_w8(AVIOContext *s, int b);
func (s *AVIOContext) AvioW8(b ffcommon.FInt) {
	return
}

//void avio_write(AVIOContext *s, const unsigned char *buf, int size);
func (s *AVIOContext) AvioWrite(buf ffcommon.FUnsignedCharP, b ffcommon.FInt) {
	return
}

//void avio_wl64(AVIOContext *s, uint64_t val);
func (s *AVIOContext) AvioWl64(val ffcommon.FInt64T) {
	return
}

//void avio_wb64(AVIOContext *s, uint64_t val);
func (s *AVIOContext) AvioWb64(val ffcommon.FInt64T) {
	return
}

//void avio_wl32(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWl32(val ffcommon.FInt) {
	return
}

//void avio_wb32(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWb32(val ffcommon.FInt) {
	return
}

//void avio_wl24(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWl24(val ffcommon.FInt) {
	return
}

//void avio_wb24(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWb24(val ffcommon.FInt) {
	return
}

//void avio_wl16(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWl16(val ffcommon.FUnsignedInt) {
	return
}

//void avio_wb16(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWb16(val ffcommon.FUnsignedInt) {
	return
}

/**
 * Write a NULL-terminated string.
 * @return number of bytes written.
 */
//int avio_put_str(AVIOContext *s, const char *str);
func (s *AVIOContext) AvioPutStr(str ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Convert an UTF-8 string to UTF-16LE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
//int avio_put_str16le(AVIOContext *s, const char *str);
func (s *AVIOContext) AvioPutStr16le(str ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Convert an UTF-8 string to UTF-16BE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
//int avio_put_str16be(AVIOContext *s, const char *str);
func (s *AVIOContext) AvioPutStr16be(str ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Mark the written bytestream as a specific type.
 *
 * Zero-length ranges are omitted from the output.
 *
 * @param time the stream time the current bytestream pos corresponds to
 *             (in AV_TIME_BASE units), or AV_NOPTS_VALUE if unknown or not
 *             applicable
 * @param type the kind of data written starting at the current pos
 */
//void avio_write_marker(AVIOContext *s, int64_t time, enum AVIODataMarkerType type);
func (s *AVIOContext) AvioWriteMarker(time ffcommon.FInt64T, type0 ffconstant.AVIODataMarkerType) {
	return
}

/**
 * fseek() equivalent for AVIOContext.
 * @return new position or AVERROR.
 */
//int64_t avio_seek(AVIOContext *s, int64_t offset, int whence);
func (s *AVIOContext) AvioSeek(offset ffcommon.FInt64T, whence ffcommon.FInt) (res ffcommon.FInt64T) {
	return
}

/**
 * Skip given number of bytes forward
 * @return new position or AVERROR.
 */
//int64_t avio_skip(AVIOContext *s, int64_t offset);
func (s *AVIOContext) AvioSkip(offset ffcommon.FInt64T) (res ffcommon.FInt64T) {
	return
}

/**
 * Get the filesize.
 * @return filesize or AVERROR
 */
//int64_t avio_size(AVIOContext *s);
func (s *AVIOContext) AvioSize() (res ffcommon.FInt64T) {
	return
}

/**
 * Similar to feof() but also returns nonzero on read errors.
 * @return non zero if and only if at end of file or a read error happened when reading.
 */
//int avio_feof(AVIOContext *s);
func (s *AVIOContext) AvioFeof() (res ffcommon.FInt) {
	return
}

/**
 * Writes a formatted string to the context.
 * @return number of bytes written, < 0 on error.
 */
//int avio_printf(AVIOContext *s, const char *fmt, ...) av_printf_format(2, 3);
func (s *AVIOContext) AvioPrintf(fmt0 ...ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Write a NULL terminated array of strings to the context.
 * Usually you don't need to use this function directly but its macro wrapper,
 * avio_print.
 */
//void avio_print_string_array(AVIOContext *s, const char *strings[]);
func (s *AVIOContext) AvioPrintStringArray(strings0 []ffcommon.FConstCharP) {
	return
}

/**
 * Force flushing of buffered data.
 *
 * For write streams, force the buffered data to be immediately written to the output,
 * without to wait to fill the internal buffer.
 *
 * For read streams, discard all currently buffered data, and advance the
 * reported file position to that of the underlying stream. This does not
 * read new data, and does not perform any seeks.
 */
//void avio_flush(AVIOContext *s);
func (s *AVIOContext) AvioFlush() {
	return
}

/**
 * Read size bytes from AVIOContext into buf.
 * @return number of bytes read or AVERROR
 */
//int avio_read(AVIOContext *s, unsigned char *buf, int size);
func (s *AVIOContext) AvioRead(buf ffcommon.FUnsignedCharP, size ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Read size bytes from AVIOContext into buf. Unlike avio_read(), this is allowed
 * to read fewer bytes than requested. The missing bytes can be read in the next
 * call. This always tries to read at least 1 byte.
 * Useful to reduce latency in certain cases.
 * @return number of bytes read or AVERROR
 */
//int avio_read_partial(AVIOContext *s, unsigned char *buf, int size);
func (s *AVIOContext) AvioReadPartial(buf ffcommon.FUnsignedCharP, size ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * @name Functions for reading from AVIOContext
 * @{
 *
 * @note return 0 if EOF, so you cannot use it if EOF handling is
 *       necessary
 */
//int          avio_r8  (AVIOContext *s);
func (s *AVIOContext) AvioR8() (res ffcommon.FInt) {
	return
}

//unsigned int avio_rl16(AVIOContext *s);
func (s *AVIOContext) AvioRl16() (res ffcommon.FUnsignedInt) {
	return
}

//unsigned int avio_rl24(AVIOContext *s);
func (s *AVIOContext) AvioRl24() (res ffcommon.FUnsignedInt) {
	return
}

//unsigned int avio_rl32(AVIOContext *s);
func (s *AVIOContext) AvioRl32() (res ffcommon.FUnsignedInt) {
	return
}

//uint64_t     avio_rl64(AVIOContext *s);
func (s *AVIOContext) AvioRl64() (res ffcommon.FUint64T) {
	return
}

//unsigned int avio_rb16(AVIOContext *s);
func (s *AVIOContext) AvioRb16() (res ffcommon.FUnsignedInt) {
	return
}

//unsigned int avio_rb24(AVIOContext *s);
func (s *AVIOContext) AvioRb24() (res ffcommon.FUnsignedInt) {
	return
}

//unsigned int avio_rb32(AVIOContext *s);
func (s *AVIOContext) AvioRb32() (res ffcommon.FUnsignedInt) {
	return
}

//uint64_t     avio_rb64(AVIOContext *s);
func (s *AVIOContext) AvioRb64() (res ffcommon.FUint64T) {
	return
}

/**
 * @}
 */

/**
 * Read a string from pb into buf. The reading will terminate when either
 * a NULL character was encountered, maxlen bytes have been read, or nothing
 * more can be read from pb. The result is guaranteed to be NULL-terminated, it
 * will be truncated if buf is too small.
 * Note that the string is not interpreted or validated in any way, it
 * might get truncated in the middle of a sequence for multi-byte encodings.
 *
 * @return number of bytes read (is always <= maxlen).
 * If reading ends on EOF or error, the return value will be one more than
 * bytes actually read.
 */
//int avio_get_str(AVIOContext *pb, int maxlen, char *buf, int buflen);
func (pb *AVIOContext) AvioGetStr(maxlen ffcommon.FInt, buf ffcommon.FConstCharP, buflen ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Read a UTF-16 string from pb and convert it to UTF-8.
 * The reading will terminate when either a null or invalid character was
 * encountered or maxlen bytes have been read.
 * @return number of bytes read (is always <= maxlen)
 */
//int avio_get_str16le(AVIOContext *pb, int maxlen, char *buf, int buflen);
func (pb *AVIOContext) AvioGetStr16le(maxlen ffcommon.FInt, buf ffcommon.FConstCharP, buflen ffcommon.FInt) (res ffcommon.FInt) {
	return
}

//int avio_get_str16be(AVIOContext *pb, int maxlen, char *buf, int buflen);
func (pb *AVIOContext) AvioGetStr16be(maxlen ffcommon.FInt, buf ffcommon.FConstCharP, buflen ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
//int avio_open(AVIOContext **s, const char *url, int flags);
func (s *AVIOContext) AvioOpen(url ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	//s取地址
	return
}

/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @param int_cb an interrupt callback to be used at the protocols level
 * @param options  A dictionary filled with protocol-private options. On return
 * this parameter will be destroyed and replaced with a dict containing options
 * that were not found. May be NULL.
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
//int avio_open2(AVIOContext **s, const char *url, int flags,
//const AVIOInterruptCB *int_cb, AVDictionary **options);
func (s *AVIOContext) AvioOpen2(url ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	//s取地址
	return
}

/**
 * Close the resource accessed by the AVIOContext s and free it.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_closep
 */
//int avio_close(AVIOContext *s);
func (s *AVIOContext) AvioClose() (res ffcommon.FInt) {
	return
}

/**
 * Close the resource accessed by the AVIOContext *s, free it
 * and set the pointer pointing to it to NULL.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_close
 */
//int avio_closep(AVIOContext **s);
func (s *AVIOContext) AvioClosep() (res ffcommon.FInt) {
	//s取地址
	return
}

/**
 * Open a write only memory stream.
 *
 * @param s new IO context
 * @return zero if no error.
 */
//int avio_open_dyn_buf(AVIOContext **s);
func (s *AVIOContext) AvioOpenDynBuf() (res ffcommon.FInt) {
	//s取地址
	return
}

/**
 * Return the written size and a pointer to the buffer.
 * The AVIOContext stream is left intact.
 * The buffer must NOT be freed.
 * No padding is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
//int avio_get_dyn_buf(AVIOContext *s, uint8_t **pbuffer);
func (s *AVIOContext) AvioGetDynBuf(pbuffer **ffcommon.FUint8T) (res ffcommon.FInt) {
	return
}

/**
 * Return the written size and a pointer to the buffer. The buffer
 * must be freed with av_free().
 * Padding of AV_INPUT_BUFFER_PADDING_SIZE is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
//int avio_close_dyn_buf(AVIOContext *s, uint8_t **pbuffer);
func (s *AVIOContext) AvioCloseDynBuf(pbuffer **ffcommon.FUint8T) (res ffcommon.FInt) {
	return
}

/**
 * Iterate through names of available protocols.
 *
 * @param opaque A private pointer representing current protocol.
 *        It must be a pointer to NULL on first iteration and will
 *        be updated by successive calls to avio_enum_protocols.
 * @param output If set to 1, iterate over output protocols,
 *               otherwise over input protocols.
 *
 * @return A static string containing the name of current protocol or NULL
 */
//const char *avio_enum_protocols(void **opaque, int output);
func AvioCloseDynBuf(opaque *ffcommon.FVoidP, output ffcommon.FInt) (res ffcommon.FConstCharP) {
	return
}

/**
 * Get AVClass by names of available protocols.
 *
 * @return A AVClass of input protocol name or NULL
 */
//const AVClass *avio_protocol_get_class(const char *name);
func AvioProtocolGetClass(name ffcommon.FConstCharP) (res *AVClass) {
	return
}

/**
 * Pause and resume playing - only meaningful if using a network streaming
 * protocol (e.g. MMS).
 *
 * @param h     IO context from which to call the read_pause function pointer
 * @param pause 1 for pause, 0 for resume
 */
//int     avio_pause(AVIOContext *h, int pause);
func (h *AVIOContext) AvioPause(pause ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Seek to a given timestamp relative to some component stream.
 * Only meaningful if using a network streaming protocol (e.g. MMS.).
 *
 * @param h IO context from which to call the seek function pointers
 * @param stream_index The stream index that the timestamp is relative to.
 *        If stream_index is (-1) the timestamp should be in AV_TIME_BASE
 *        units from the beginning of the presentation.
 *        If a stream_index >= 0 is used and the protocol does not support
 *        seeking based on component streams, the call will fail.
 * @param timestamp timestamp in AVStream.time_base units
 *        or if there is no stream specified then in AV_TIME_BASE units.
 * @param flags Optional combination of AVSEEK_FLAG_BACKWARD, AVSEEK_FLAG_BYTE
 *        and AVSEEK_FLAG_ANY. The protocol may silently ignore
 *        AVSEEK_FLAG_BACKWARD and AVSEEK_FLAG_ANY, but AVSEEK_FLAG_BYTE will
 *        fail if used and not supported.
 * @return >= 0 on success
 * @see AVInputFormat::read_seek
 */
//int64_t avio_seek_time(AVIOContext *h, int stream_index,
//int64_t timestamp, int flags);
func (h *AVIOContext) AvioSeekTime(stream_index ffcommon.FInt,
	timestamp ffcommon.FInt64T, flags ffcommon.FInt) (res ffcommon.FInt64T) {
	return
}

/* Avoid a warning. The header can not be included because it breaks c++. */
type AVBPrint struct {
}

/**
 * Read contents of h into print buffer, up to max_size bytes, or up to EOF.
 *
 * @return 0 for success (max_size bytes read or EOF reached), negative error
 * code otherwise
 */
//int avio_read_to_bprint(AVIOContext *h, struct AVBPrint *pb, size_t max_size);
func (h *AVIOContext) AvioReadToBprint(pb *AVBPrint, max_size ffcommon.FSizeT) (res ffcommon.FInt) {
	return
}

/**
 * Accept and allocate a client context on a server context.
 * @param  s the server context
 * @param  c the client context, must be unallocated
 * @return   >= 0 on success or a negative value corresponding
 *           to an AVERROR on failure
 */
//int avio_accept(AVIOContext *s, AVIOContext **c);
func (s *AVIOContext) AvioAccept(c **AVIOContext) (res ffcommon.FInt) {
	return
}

/**
 * Perform one step of the protocol handshake to accept a new client.
 * This function must be called on a client returned by avio_accept() before
 * using it as a read/write context.
 * It is separate from avio_accept() because it may block.
 * A step of the handshake is defined by places where the application may
 * decide to change the proceedings.
 * For example, on a protocol with a request header and a reply header, each
 * one can constitute a step because the application may use the parameters
 * from the request to change parameters in the reply; or each individual
 * chunk of the request can constitute a step.
 * If the handshake is already finished, avio_handshake() does nothing and
 * returns 0 immediately.
 *
 * @param  c the client context to perform the handshake on
 * @return   0   on a complete and successful handshake
 *           > 0 if the handshake progressed, but is not complete
 *           < 0 for an AVERROR code
 */
//int avio_handshake(AVIOContext *c);
func (c *AVIOContext) AvioHandshake() (res ffcommon.FInt) {
	return
}
