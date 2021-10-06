package libswresample

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
)

/**
 * The libswresample context. Unlike libavcodec and libavformat, this structure
 * is opaque. This means that if you would like to set options, you must use
 * the @ref avoptions API and cannot directly set values to members of the
 * structure.
 */
type SwrContext struct {
}

/**
 * Get the AVClass for SwrContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 * @return the AVClass of SwrContext
 */
//const AVClass *swr_get_class(void);
func SwrGetClass() (res *AVClass) {
	return
}

/**
 * @name SwrContext constructor functions
 * @{
 */

/**
 * Allocate SwrContext.
 *
 * If you use this function you will need to set the parameters (manually or
 * with swr_alloc_set_opts()) before calling swr_init().
 *
 * @see swr_alloc_set_opts(), swr_init(), swr_free()
 * @return NULL on error, allocated context otherwise
 */
//struct SwrContext *swr_alloc(void);
func SwrAlloc() (res *SwrContext) {
	return
}

/**
 * Initialize context after user parameters have been set.
 * @note The context must be configured using the AVOption API.
 *
 * @see av_opt_set_int()
 * @see av_opt_set_dict()
 *
 * @param[in,out]   s Swr context to initialize
 * @return AVERROR error code in case of failure.
 */
//int swr_init(struct SwrContext *s);
func (s *SwrContext) SwrInit() (res ffcommon.FInt) {
	return
}

/**
 * Check whether an swr context has been initialized or not.
 *
 * @param[in]       s Swr context to check
 * @see swr_init()
 * @return positive if it has been initialized, 0 if not initialized
 */
//int swr_is_initialized(struct SwrContext *s);
func (s *SwrContext) SwrIsInitialized() (res ffcommon.FInt) {
	return
}

/**
 * Allocate SwrContext if needed and set/reset common parameters.
 *
 * This function does not require s to be allocated with swr_alloc(). On the
 * other hand, swr_alloc() can use swr_alloc_set_opts() to set the parameters
 * on the allocated context.
 *
 * @param s               existing Swr context if available, or NULL if not
 * @param out_ch_layout   output channel layout (AV_CH_LAYOUT_*)
 * @param out_sample_fmt  output sample format (AV_SAMPLE_FMT_*).
 * @param out_sample_rate output sample rate (frequency in Hz)
 * @param in_ch_layout    input channel layout (AV_CH_LAYOUT_*)
 * @param in_sample_fmt   input sample format (AV_SAMPLE_FMT_*).
 * @param in_sample_rate  input sample rate (frequency in Hz)
 * @param log_offset      logging level offset
 * @param log_ctx         parent logging context, can be NULL
 *
 * @see swr_init(), swr_free()
 * @return NULL on error, allocated context otherwise
 */
//struct SwrContext *swr_alloc_set_opts(struct SwrContext *s,
//int64_t out_ch_layout, enum AVSampleFormat out_sample_fmt, int out_sample_rate,
//int64_t  in_ch_layout, enum AVSampleFormat  in_sample_fmt, int  in_sample_rate,
//int log_offset, void *log_ctx);
func (s *SwrContext) swr_alloc_set_opts(out_ch_layout ffcommon.FInt64T, out_sample_fmt ffconstant.AVSampleFormat, out_sample_rate ffcommon.FInt,
	in_ch_layout ffcommon.FInt64T, in_sample_fmt ffconstant.AVSampleFormat, in_sample_rate ffcommon.FInt,
	log_offset ffcommon.FInt, log_ctx ffcommon.FVoidP) (res *SwrContext) {
	return
}

/**
 * @}
 *
 * @name SwrContext destructor functions
 * @{
 */

/**
 * Free the given SwrContext and set the pointer to NULL.
 *
 * @param[in] s a pointer to a pointer to Swr context
 */
//void swr_free(struct SwrContext **s);
func (s *SwrContext) SwrFree() {
	//s再取地址
	return
}

/**
 * Closes the context so that swr_is_initialized() returns 0.
 *
 * The context can be brought back to life by running swr_init(),
 * swr_init() can also be used without swr_close().
 * This function is mainly provided for simplifying the usecase
 * where one tries to support libavresample and libswresample.
 *
 * @param[in,out] s Swr context to be closed
 */
//void swr_close(struct SwrContext *s);
func (s *SwrContext) SwrClose() {
	return
}

/**
 * @}
 *
 * @name Core conversion functions
 * @{
 */

/** Convert audio.
 *
 * in and in_count can be set to 0 to flush the last few samples out at the
 * end.
 *
 * If more input is provided than output space, then the input will be buffered.
 * You can avoid this buffering by using swr_get_out_samples() to retrieve an
 * upper bound on the required number of output samples for the given number of
 * input samples. Conversion will run directly without copying whenever possible.
 *
 * @param s         allocated Swr context, with parameters set
 * @param out       output buffers, only the first one need be set in case of packed audio
 * @param out_count amount of space available for output in samples per channel
 * @param in        input buffers, only the first one need to be set in case of packed audio
 * @param in_count  number of input samples available in one channel
 *
 * @return number of samples output per channel, negative value on error
 */
//int swr_convert(struct SwrContext *s, uint8_t **out, int out_count,
//const uint8_t **in , int in_count);
func (s *SwrContext) SwrConvert(out **ffcommon.FUint8T, out_count ffcommon.FInt,
	in **ffcommon.FUint8T, in_count ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Convert the next timestamp from input to output
 * timestamps are in 1/(in_sample_rate * out_sample_rate) units.
 *
 * @note There are 2 slightly differently behaving modes.
 *       @li When automatic timestamp compensation is not used, (min_compensation >= FLT_MAX)
 *              in this case timestamps will be passed through with delays compensated
 *       @li When automatic timestamp compensation is used, (min_compensation < FLT_MAX)
 *              in this case the output timestamps will match output sample numbers.
 *              See ffmpeg-resampler(1) for the two modes of compensation.
 *
 * @param s[in]     initialized Swr context
 * @param pts[in]   timestamp for the next input sample, INT64_MIN if unknown
 * @see swr_set_compensation(), swr_drop_output(), and swr_inject_silence() are
 *      function used internally for timestamp compensation.
 * @return the output timestamp for the next output sample
 */
//int64_t swr_next_pts(struct SwrContext *s, int64_t pts);
func (s *SwrContext) SwrNextPts(pts ffcommon.FInt64T) (res ffcommon.FInt64T) {
	return
}

/**
 * @}
 *
 * @name Low-level option setting functions
 * These functons provide a means to set low-level options that is not possible
 * with the AVOption API.
 * @{
 */

/**
 * Activate resampling compensation ("soft" compensation). This function is
 * internally called when needed in swr_next_pts().
 *
 * @param[in,out] s             allocated Swr context. If it is not initialized,
 *                              or SWR_FLAG_RESAMPLE is not set, swr_init() is
 *                              called with the flag set.
 * @param[in]     sample_delta  delta in PTS per sample
 * @param[in]     compensation_distance number of samples to compensate for
 * @return    >= 0 on success, AVERROR error codes if:
 *            @li @c s is NULL,
 *            @li @c compensation_distance is less than 0,
 *            @li @c compensation_distance is 0 but sample_delta is not,
 *            @li compensation unsupported by resampler, or
 *            @li swr_init() fails when called.
 */
//int swr_set_compensation(struct SwrContext *s, int sample_delta, int compensation_distance);
func (s *SwrContext) SwrSetCompensation(sample_delta ffcommon.FInt, compensation_distance ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Set a customized input channel mapping.
 *
 * @param[in,out] s           allocated Swr context, not yet initialized
 * @param[in]     channel_map customized input channel mapping (array of channel
 *                            indexes, -1 for a muted channel)
 * @return >= 0 on success, or AVERROR error code in case of failure.
 */
//int swr_set_channel_mapping(struct SwrContext *s, const int *channel_map);
func (s *SwrContext) SwrSetChannelMapping(channel_map *ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Generate a channel mixing matrix.
 *
 * This function is the one used internally by libswresample for building the
 * default mixing matrix. It is made public just as a utility function for
 * building custom matrices.
 *
 * @param in_layout           input channel layout
 * @param out_layout          output channel layout
 * @param center_mix_level    mix level for the center channel
 * @param surround_mix_level  mix level for the surround channel(s)
 * @param lfe_mix_level       mix level for the low-frequency effects channel
 * @param rematrix_maxval     if 1.0, coefficients will be normalized to prevent
 *                            overflow. if INT_MAX, coefficients will not be
 *                            normalized.
 * @param[out] matrix         mixing coefficients; matrix[i + stride * o] is
 *                            the weight of input channel i in output channel o.
 * @param stride              distance between adjacent input channels in the
 *                            matrix array
 * @param matrix_encoding     matrixed stereo downmix mode (e.g. dplii)
 * @param log_ctx             parent logging context, can be NULL
 * @return                    0 on success, negative AVERROR code on failure
 */
//int swr_build_matrix(uint64_t in_layout, uint64_t out_layout,
//double center_mix_level, double surround_mix_level,
//double lfe_mix_level, double rematrix_maxval,
//double rematrix_volume, double *matrix,
//int stride, enum AVMatrixEncoding matrix_encoding,
//void *log_ctx);
func SwrBuildMatrix(in_layout ffcommon.FUint64T, out_layout ffcommon.FUint64T,
	center_mix_level ffcommon.FDouble, surround_mix_level ffcommon.FDouble,
	lfe_mix_level ffcommon.FDouble, rematrix_maxval ffcommon.FDouble,
	rematrix_volume ffcommon.FDouble, matrix *ffcommon.FDouble,
	stride ffcommon.FInt, matrix_encoding ffconstant.AVMatrixEncoding,
	log_ctx ffcommon.FVoidP) (res ffcommon.FInt) {
	return
}

/**
 * Set a customized remix matrix.
 *
 * @param s       allocated Swr context, not yet initialized
 * @param matrix  remix coefficients; matrix[i + stride * o] is
 *                the weight of input channel i in output channel o
 * @param stride  offset between lines of the matrix
 * @return  >= 0 on success, or AVERROR error code in case of failure.
 */
//int swr_set_matrix(struct SwrContext *s, const double *matrix, int stride);
func (s *SwrContext) SwrSetMatrix(matrix *ffcommon.FDouble, stride ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * @}
 *
 * @name Sample handling functions
 * @{
 */

/**
 * Drops the specified number of output samples.
 *
 * This function, along with swr_inject_silence(), is called by swr_next_pts()
 * if needed for "hard" compensation.
 *
 * @param s     allocated Swr context
 * @param count number of samples to be dropped
 *
 * @return >= 0 on success, or a negative AVERROR code on failure
 */
//int swr_drop_output(struct SwrContext *s, int count);
func (s *SwrContext) SwrDropOutput(count ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Injects the specified number of silence samples.
 *
 * This function, along with swr_drop_output(), is called by swr_next_pts()
 * if needed for "hard" compensation.
 *
 * @param s     allocated Swr context
 * @param count number of samples to be dropped
 *
 * @return >= 0 on success, or a negative AVERROR code on failure
 */
//int swr_inject_silence(struct SwrContext *s, int count);
func (s *SwrContext) SwrInjectSilence(count ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Gets the delay the next input sample will experience relative to the next output sample.
 *
 * Swresample can buffer data if more input has been provided than available
 * output space, also converting between sample rates needs a delay.
 * This function returns the sum of all such delays.
 * The exact delay is not necessarily an integer value in either input or
 * output sample rate. Especially when downsampling by a large value, the
 * output sample rate may be a poor choice to represent the delay, similarly
 * for upsampling and the input sample rate.
 *
 * @param s     swr context
 * @param base  timebase in which the returned delay will be:
 *              @li if it's set to 1 the returned delay is in seconds
 *              @li if it's set to 1000 the returned delay is in milliseconds
 *              @li if it's set to the input sample rate then the returned
 *                  delay is in input samples
 *              @li if it's set to the output sample rate then the returned
 *                  delay is in output samples
 *              @li if it's the least common multiple of in_sample_rate and
 *                  out_sample_rate then an exact rounding-free delay will be
 *                  returned
 * @returns     the delay in 1 / @c base units.
 */
//int64_t swr_get_delay(struct SwrContext *s, int64_t base);
func (s *SwrContext) SwrGetDelay(base ffcommon.FInt64T) (res ffcommon.FInt64T) {
	return
}

/**
 * Find an upper bound on the number of samples that the next swr_convert
 * call will output, if called with in_samples of input samples. This
 * depends on the internal state, and anything changing the internal state
 * (like further swr_convert() calls) will may change the number of samples
 * swr_get_out_samples() returns for the same number of input samples.
 *
 * @param in_samples    number of input samples.
 * @note any call to swr_inject_silence(), swr_convert(), swr_next_pts()
 *       or swr_set_compensation() invalidates this limit
 * @note it is recommended to pass the correct available buffer size
 *       to all functions like swr_convert() even if swr_get_out_samples()
 *       indicates that less would be used.
 * @returns an upper bound on the number of samples that the next swr_convert
 *          will output or a negative value to indicate an error
 */
//int swr_get_out_samples(struct SwrContext *s, int in_samples);
func (s *SwrContext) SwrGetOutSamples(in_samples ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * @}
 *
 * @name Configuration accessors
 * @{
 */

/**
 * Return the @ref LIBSWRESAMPLE_VERSION_INT ffconstant.
 *
 * This is useful to check if the build-time libswresample has the same version
 * as the run-time one.
 *
 * @returns     the unsigned int-typed version
 */
//unsigned swresample_version(void);
func SwresampleVersion() (res ffcommon.FUnsigned) {
	return
}

/**
 * Return the swr build-time configuration.
 *
 * @returns     the build-time @c ./configure flags
 */
//const char *swresample_configuration(void);
func SwresampleConfiguration() (res ffcommon.FConstCharP) {
	return
}

/**
 * Return the swr license.
 *
 * @returns     the license of libswresample, determined at build-time
 */
//const char *swresample_license(void);
func SwresampleLicense() (res ffcommon.FConstCharP) {
	return
}

/**
 * @}
 *
 * @name AVFrame based API
 * @{
 */

/**
 * Convert the samples in the input AVFrame and write them to the output AVFrame.
 *
 * Input and output AVFrames must have channel_layout, sample_rate and format set.
 *
 * If the output AVFrame does not have the data pointers allocated the nb_samples
 * field will be set using av_frame_get_buffer()
 * is called to allocate the frame.
 *
 * The output AVFrame can be NULL or have fewer allocated samples than required.
 * In this case, any remaining samples not written to the output will be added
 * to an internal FIFO buffer, to be returned at the next call to this function
 * or to swr_convert().
 *
 * If converting sample rate, there may be data remaining in the internal
 * resampling delay buffer. swr_get_delay() tells the number of
 * remaining samples. To get this data as output, call this function or
 * swr_convert() with NULL input.
 *
 * If the SwrContext configuration does not match the output and
 * input AVFrame settings the conversion does not take place and depending on
 * which AVFrame is not matching AVERROR_OUTPUT_CHANGED, AVERROR_INPUT_CHANGED
 * or the result of a bitwise-OR of them is returned.
 *
 * @see swr_delay()
 * @see swr_convert()
 * @see swr_get_delay()
 *
 * @param swr             audio resample context
 * @param output          output AVFrame
 * @param input           input AVFrame
 * @return                0 on success, AVERROR on failure or nonmatching
 *                        configuration.
 */
//int swr_convert_frame(SwrContext *swr,
//AVFrame *output, const AVFrame *input);
func (swr *SwrContext) SwrConvertFrame(output *AVFrame, input *AVFrame) (res ffcommon.FInt) {
	return
}

/**
 * Configure or reconfigure the SwrContext using the information
 * provided by the AVFrames.
 *
 * The original resampling context is reset even on failure.
 * The function calls swr_close() internally if the context is open.
 *
 * @see swr_close();
 *
 * @param swr             audio resample context
 * @param output          output AVFrame
 * @param input           input AVFrame
 * @return                0 on success, AVERROR on failure.
 */
//int swr_config_frame(SwrContext *swr, const AVFrame *out, const AVFrame *in);
func (swr *SwrContext) SwrConfigFrame(out *AVFrame, in *AVFrame) (res ffcommon.FInt) {
	return
}
