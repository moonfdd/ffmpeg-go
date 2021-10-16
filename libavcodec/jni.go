package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Manually set a Java virtual machine which will be used to retrieve the JNI
 * environment. Once a Java VM is set it cannot be changed afterwards, meaning
 * you can call multiple times av_jni_set_java_vm with the same Java VM pointer
 * however it will error out if you try to set a different Java VM.
 *
 * @param vm Java virtual machine
 * @param log_ctx context used for logging, can be NULL
 * @return 0 on success, < 0 otherwise
 */
//int av_jni_set_java_vm(void *vm, void *log_ctx);
//未测试
func AvJniSetJavaVm(vm ffcommon.FVoidP, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_jni_set_java_vm").Call(
		uintptr(unsafe.Pointer(vm)),
		uintptr(unsafe.Pointer(log_ctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/*
 * Get the Java virtual machine which has been set with av_jni_set_java_vm.
 *
 * @param vm Java virtual machine
 * @return a pointer to the Java virtual machine
 */
//void *av_jni_get_java_vm(void *log_ctx);
//未测试
func AvJniGetJavaVm(log_ctx ffcommon.FVoidP) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_jni_get_java_vm").Call(
		uintptr(unsafe.Pointer(log_ctx)),
	)
	if err != nil {
		//return
	}
	res = t
	return
}
