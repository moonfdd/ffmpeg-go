package ffconstant

/* values for the flags, the stuff on the command line is different */
const SWS_FAST_BILINEAR = 1
const SWS_BILINEAR = 2
const SWS_BICUBIC = 4
const SWS_X = 8
const SWS_POINT = 0x10
const SWS_AREA = 0x20
const SWS_BICUBLIN = 0x40
const SWS_GAUSS = 0x80
const SWS_SINC = 0x100
const SWS_LANCZOS = 0x200
const SWS_SPLINE = 0x400

const SWS_SRC_V_CHR_DROP_MASK = 0x30000
const SWS_SRC_V_CHR_DROP_SHIFT = 16

const SWS_PARAM_DEFAULT = 123456

const SWS_PRINT_INFO = 0x1000

//the following 3 flags are not completely implemented
//internal chrominance subsampling info
const SWS_FULL_CHR_H_INT = 0x2000

//input subsampling info
const SWS_FULL_CHR_H_INP = 0x4000
const SWS_DIRECT_BGR = 0x8000
const SWS_ACCURATE_RND = 0x40000
const SWS_BITEXACT = 0x80000
const SWS_ERROR_DIFFUSION = 0x800000

const SWS_MAX_REDUCE_CUTOFF = 0.002

const SWS_CS_ITU709 = 1
const SWS_CS_FCC = 4
const SWS_CS_ITU601 = 5
const SWS_CS_ITU624 = 5
const SWS_CS_SMPTE170M = 5
const SWS_CS_SMPTE240M = 7
const SWS_CS_DEFAULT = 5
const SWS_CS_BT2020 = 9
