package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

// go run .\examples\internalexamples\avio_reading失败\main.go D:\mysetup\gopath\src\dsy\sf3_c\x64\Debug\1.mp4
func main() {

	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvutilPath("./lib/avutil-56.dll")
	ffcommon.SetAvcodecPath("./lib/avcodec-58.dll")
	ffcommon.SetAvdevicePath("./lib/avdevice-58.dll")
	ffcommon.SetAvfilterPath("./lib/avfilter-56.dll")
	ffcommon.SetAvformatPath("./lib/avformat-58.dll")
	ffcommon.SetAvpostprocPath("./lib/postproc-55.dll")
	ffcommon.SetAvswresamplePath("./lib/swresample-3.dll")
	ffcommon.SetAvswscalePath("./lib/swscale-5.dll")

	// os.Setenv("Path", os.Getenv("Path")+";./lib7")
	// ffcommon.SetAvutilPath("./lib7/avutil-59.dll")
	// ffcommon.SetAvcodecPath("./lib7/avcodec-61.dll")
	// ffcommon.SetAvdevicePath("./lib7/avdevice-61.dll")
	// ffcommon.SetAvfilterPath("./lib7/avfilter-10.dll")
	// ffcommon.SetAvformatPath("./lib7/avformat-61.dll")
	// ffcommon.SetAvpostprocPath("./lib7/postproc-58.dll")
	// ffcommon.SetAvswresamplePath("./lib7/swresample-5.dll")
	// ffcommon.SetAvswscalePath("./lib7/swscale-8.dll")

	genDir := "./out"
	_, err := os.Stat(genDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(genDir, 0777) //  Everyone can read write and execute
		}
	}
	// return
	main0()
}

var bd *buffer_data = &buffer_data{}

func main0() (ret ffcommon.FInt) {
	// libavformat.AvRegisterAll()
	var fmt_ctx *libavformat.AVFormatContext
	var avio_ctx *libavformat.AVIOContext
	var buffer *byte
	var avio_ctx_buffer *byte
	var buffer_size uint64
	avio_ctx_buffer_size := uint64(4096)
	var input_filename string
	// bd := &buffer_data{}
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s input_file\nAPI example program to show how to read from a custom buffer accessed through AVIOContext.\n", os.Args[0])
		return 1
	}
	input_filename = os.Args[1]

	/* slurp file content into buffer */
	ret = libavutil.AvFileMap(input_filename, &buffer, &buffer_size, 0, uintptr(0))
	if ret < 0 {
		goto end
	}

	/* fill opaque structure used by the AVIOContext read callback */
	bd.Ptr = buffer
	bd.Size = buffer_size

	fmt_ctx = libavformat.AvformatAllocContext()
	if fmt_ctx == nil {
		ret = -libavutil.ENOMEM
		goto end
	}

	avio_ctx_buffer = (*byte)(unsafe.Pointer(libavutil.AvMalloc(avio_ctx_buffer_size)))
	fmt.Printf("avio_ctx_buffer = %d\r\n", avio_ctx_buffer)
	if avio_ctx_buffer == nil {
		ret = -libavutil.ENOMEM
		goto end
	}
	avio_ctx = libavformat.AvioAllocContext(avio_ctx_buffer, int32(avio_ctx_buffer_size),
		0, uintptr(unsafe.Pointer(bd)), read_packet, nil, nil)
	if avio_ctx == nil {
		ret = -libavutil.ENOMEM
		goto end
	}

	fmt_ctx.Pb = avio_ctx

	ret = libavformat.AvformatOpenInput(&fmt_ctx, "", nil, nil)
	if ret < 0 {
		fmt.Printf("Could not open input %d\n", ret)
		goto end
	}
	ret = fmt_ctx.AvformatFindStreamInfo(nil)
	if ret < 0 {
		fmt.Printf("Could not find stream information\n")
		goto end
	}

	fmt_ctx.AvDumpFormat(0, input_filename, 0)

end:
	libavformat.AvformatCloseInput(&fmt_ctx)

	/* note: the internal buffer could have changed, and be != avio_ctx_buffer */
	if avio_ctx != nil {
		libavutil.AvFreep(uintptr(unsafe.Pointer(&avio_ctx.Buffer)))
	}
	libavformat.AvioContextFree(&avio_ctx)

	libavutil.AvFileUnmap(buffer, buffer_size)

	if ret < 0 {
		fmt.Printf("Error occurred: %s\n", libavutil.AvErr2str(ret))
		return 1
	}

	return 0
}

type buffer_data struct {
	Ptr  *byte
	Size ffcommon.FSizeT ///< size left in the buffer
}

func read_packet(opaque uintptr, buf *byte, buf_size int32) uintptr {
	fmt.Printf("opaque:%d buf:%d buf_size=%d\n", opaque, buf, buf_size)
	// bd := (*buffer_data)(unsafe.Pointer(opaque))
	if buf_size > int32(bd.Size) {
		buf_size = int32(bd.Size)
	}
	fmt.Printf("bd.Ptr:%d  bd.Size=%d buf_size=%d\n", bd.Ptr, bd.Size, buf_size)
	if buf_size == 0 {
		r := int32(libavutil.AVERROR_EOF)
		return uintptr(r)
	}
	// fmt.Println("buf = ", uintptr(unsafe.Pointer(buf)))

	// a := make([]byte, buf_size)
	// buf = (*byte)(unsafe.Pointer(&a[0]))
	// ccc := ffcommon.ByteSliceFromByteP(bd.Ptr, int(buf_size))[int(buf_size)-1]
	// _ = ccc
	// fmt.Println("ccc = ", ccc)
	copy(ffcommon.ByteSliceFromByteP(buf, int(buf_size)), ffcommon.ByteSliceFromByteP(bd.Ptr, int(buf_size)))

	// /* copy internal buffer data to buf */
	// for i := int32(0); i < buf_size; i++ {
	// 	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(i))) = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(bd.Ptr)) + uintptr(i)))
	// }
	bd.Ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(bd.Ptr)) + uintptr(buf_size)))
	bd.Size -= ffcommon.FSizeT(buf_size)
	return *(*uintptr)(unsafe.Pointer(&buf_size))
}

func write_packet(opaque uintptr, buf *byte, buf_size int32) uintptr {
	fmt.Println("write_packet = ", buf_size)
	return uintptr(buf_size)
}
