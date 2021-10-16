package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVAESCTR struct {
}

/**
 * Allocate an AVAESCTR context.
 */
//struct AVAESCTR *av_aes_ctr_alloc(void);
//未测试
func AvAesCtrAlloc() (res *AVAESCTR, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVAESCTR)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
 * Initialize an AVAESCTR context.
 * @param key encryption key, must have a length of AES_CTR_KEY_SIZE
 */
//int av_aes_ctr_init(struct AVAESCTR *a, const uint8_t *key);
//未测试
func (d *AVAESCTR) AvAesCtrInit(key *ffcommon.FUint8T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_init").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(key)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Release an AVAESCTR context.
 */
//void av_aes_ctr_free(struct AVAESCTR *a);
//未测试
func (a *AVDES) AvAesCtrFree() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_free").Call(
		uintptr(unsafe.Pointer(a)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Process a buffer using a previously initialized context.
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param size the size of src and dst
 */
//void av_aes_ctr_crypt(struct AVAESCTR *a, uint8_t *dst, const uint8_t *src, int size);
//未测试
func (d *AVAESCTR) AvAesCtrCrypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt) (err error) {
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_crypt").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
	)
	if err != nil {
		//return
	}
	return
}

/**
 * Get the current iv
 */
//const uint8_t* av_aes_ctr_get_iv(struct AVAESCTR *a);
//未测试
func (d *AVAESCTR) AvAesCtrGetIv() (res *ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_get_iv").Call(
		uintptr(unsafe.Pointer(d)),
	)
	if err != nil {
		//return
	}
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}

/**
 * Generate a random iv
 */
//void av_aes_ctr_set_random_iv(struct AVAESCTR *a);
//未测试
func (d *AVAESCTR) AvAesCtrSetRandomIv() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_set_random_iv").Call(
		uintptr(unsafe.Pointer(d)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Forcefully change the 8-byte iv
 */
//void av_aes_ctr_set_iv(struct AVAESCTR *a, const uint8_t* iv);
//未测试
func (d *AVAESCTR) AvAesCtrSetIv(iv *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_set_iv").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(iv)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Forcefully change the "full" 16-byte iv, including the counter
 */
//void av_aes_ctr_set_full_iv(struct AVAESCTR *a, const uint8_t* iv);
//未测试
func (d *AVAESCTR) AvAesCtrSetFullIv(iv *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_set_full_iv").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(iv)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Increment the top 64 bit of the iv (performed after each frame)
 */
//void av_aes_ctr_increment_iv(struct AVAESCTR *a);
//未测试
func (d *AVAESCTR) AvAesCtrIncrementIv() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_ctr_increment_iv").Call(
		uintptr(unsafe.Pointer(d)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
