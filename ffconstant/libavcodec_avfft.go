package ffconstant

/* Real Discrete Fourier Transform */
type RDFTransformType int32

const (
	DFT_R2C = 0
	IDFT_C2R
	IDFT_R2C
	DFT_C2R
)

type DCTTransformType int32

const (
	DCT_II = 0
	DCT_III
	DCT_I
	DST_I
)
