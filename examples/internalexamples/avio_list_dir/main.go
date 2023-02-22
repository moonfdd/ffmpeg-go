package main

import (
	"fmt"
	"os"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

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

	genDir := "./out"
	_, err := os.Stat(genDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(genDir, 0777) //  Everyone can read write and execute
		}
	}
	main0()
}
func main0() (ret ffcommon.FInt) {

	libavutil.AvLogSetLevel(libavutil.AV_LOG_DEBUG)

	if len(os.Args) < 2 {
		usage(os.Args[0])
		ret = 1
		return
	}

	libavformat.AvformatNetworkInit()

	ret = list_op(os.Args[1])

	libavformat.AvformatNetworkDeinit()

	if ret < 0 {
		ret = 1
	}
	return
}

func type_string(type0 int32) string {
	switch type0 {
	case libavformat.AVIO_ENTRY_DIRECTORY:
		return "<DIR>"
	case libavformat.AVIO_ENTRY_FILE:
		return "<FILE>"
	case libavformat.AVIO_ENTRY_BLOCK_DEVICE:
		return "<BLOCK DEVICE>"
	case libavformat.AVIO_ENTRY_CHARACTER_DEVICE:
		return "<CHARACTER DEVICE>"
	case libavformat.AVIO_ENTRY_NAMED_PIPE:
		return "<PIPE>"
	case libavformat.AVIO_ENTRY_SYMBOLIC_LINK:
		return "<LINK>"
	case libavformat.AVIO_ENTRY_SOCKET:
		return "<SOCKET>"
	case libavformat.AVIO_ENTRY_SERVER:
		return "<SERVER>"
	case libavformat.AVIO_ENTRY_SHARE:
		return "<SHARE>"
	case libavformat.AVIO_ENTRY_WORKGROUP:
		return "<WORKGROUP>"
	case libavformat.AVIO_ENTRY_UNKNOWN:
	default:
		break
	}
	return "<UNKNOWN>"
}

func list_op(input_dir string) (ret int32) {
	var entry *libavformat.AVIODirEntry
	var ctx *libavformat.AVIODirContext
	var cnt int32
	var filemode [4]byte
	var uid_and_gid [20]byte

	//注意Windows下会返回-40，也就是Function not implement，方法未实现，也就是说windows下不支持此方法
	ret = libavformat.AvioOpenDir(&ctx, input_dir, nil)
	defer libavformat.AvioCloseDir(&ctx)

	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot open directory: %s.\n", libavutil.AvErr2str(ret))
		return
	}

	cnt = 0
	for {
		ret = ctx.AvioReadDir(&entry)
		if ret < 0 {
			libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot list directory: %s.\n", libavutil.AvErr2str(ret))
			return
		}
		if entry == nil {
			break
		}
		if entry.Filemode == -1 {
			filemode[0] = '?'
			filemode[1] = '?'
			filemode[2] = '?'
		} else {
			f := fmt.Sprint(entry.Filemode)
			if len(f) >= 1 {
				filemode[0] = f[0]
			}
			if len(f) >= 2 {
				filemode[1] = f[1]
			}
			if len(f) >= 3 {
				filemode[2] = f[2]
			}
		}
		u := fmt.Sprintf("%d%d", entry.UserId, entry.GroupId)
		copy(uid_and_gid[0:19], []byte(u))

		if cnt == 0 {
			libavutil.AvLog(uintptr(0), libavutil.AV_LOG_INFO, "%-9s %12s %30s %10s %s %16s %16s %16s\n",
				"TYPE", "SIZE", "NAME", "UID(GID)", "UGO", "MODIFIED",
				"ACCESSED", "STATUS_CHANGED")
		}
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_INFO, "%-9s %12s %30s %10s %s %16s %16s %16s\n",
			type_string(entry.Type),
			fmt.Sprint(entry.Size),
			ffcommon.StringFromPtr(entry.Name),
			string(uid_and_gid[:]),
			fmt.Sprint(filemode),
			fmt.Sprint(entry.ModificationTimestamp),
			fmt.Sprint(entry.AccessTimestamp),
			fmt.Sprint(entry.StatusChangeTimestamp))
		libavformat.AvioFreeDirectoryEntry(&entry)
		cnt++
	}
	return
}

func usage(program_name string) {
	fmt.Printf("usage: %s input_dir\nAPI example program to show how to list files in directory accessed through AVIOContext.\n", program_name)
}
