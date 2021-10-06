package ffconstant

/**
 * The spec limits the number of wavelet decompositions to 4 for both
 * level 1 (VC-2) and 128 (long-gop default).
 * 5 decompositions is the maximum before >16-bit buffers are needed.
 * Schroedinger allows this for DD 9,7 and 13,7 wavelets only, limiting
 * the others to 4 decompositions (or 3 for the fidelity filter).
 *
 * We use this instead of MAX_DECOMPOSITIONS to save some memory.
 */
const MAX_DWT_LEVELS = 5

/**
 * Parse code values:
 *
 * Dirac Specification ->
 * 9.6.1  Table 9.1
 *
 * VC-2 Specification  ->
 * 10.4.1 Table 10.1
 */
type DiracParseCodes int32

const (
	DIRAC_PCODE_SEQ_HEADER      = 0x00
	DIRAC_PCODE_END_SEQ         = 0x10
	DIRAC_PCODE_AUX             = 0x20
	DIRAC_PCODE_PAD             = 0x30
	DIRAC_PCODE_PICTURE_CODED   = 0x08
	DIRAC_PCODE_PICTURE_RAW     = 0x48
	DIRAC_PCODE_PICTURE_LOW_DEL = 0xC8
	DIRAC_PCODE_PICTURE_HQ      = 0xE8
	DIRAC_PCODE_INTER_NOREF_CO1 = 0x0A
	DIRAC_PCODE_INTER_NOREF_CO2 = 0x09
	DIRAC_PCODE_INTER_REF_CO1   = 0x0D
	DIRAC_PCODE_INTER_REF_CO2   = 0x0E
	DIRAC_PCODE_INTRA_REF_CO    = 0x0C
	DIRAC_PCODE_INTRA_REF_RAW   = 0x4C
	DIRAC_PCODE_INTRA_REF_PICT  = 0xCC
	DIRAC_PCODE_MAGIC           = 0x42424344
)
