package xastiav

import (
	"github.com/asticode/go-astiav"
	"github.com/xaionaro-go/avcommon"
	xastiav "github.com/xaionaro-go/avcommon/astiav"
)

/*
// See file libavformat/libsrt.c in the ffmpeg source code
typedef struct SRTContext {
    const void *class;
    int fd;
       //...
} SRTContext;
*/
import "C"

func GetFDFromFormatContext(fmtCtx *astiav.FormatContext) C.int {
	f := avcommon.WrapAVFormatContext(xastiav.CFromAVFormatContext(fmtCtx))
	avioCtx := f.Pb()
	urlCtx := avcommon.WrapURLContext(avioCtx.Opaque())
	srtCtx := (*C.SRTContext)(urlCtx.PrivData().UnsafePointer())
	return srtCtx.fd
}
