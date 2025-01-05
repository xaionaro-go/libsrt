package xastiav

import (
	"reflect"
	"unsafe"

	"github.com/asticode/go-astiav"
	"github.com/xaionaro-go/unsafetools"
)

/*
typedef struct URLContext {
    const void *av_class;
    const struct URLProtocol *prot;
    void *priv_data;
	//...
} URLContext;

// See file libavformat/libsrt.c in the ffmpeg source code
typedef struct SRTContext {
    const void *class;
    int fd;
	//...
} SRTContext;
*/
//#include <libavcodec/avcodec.h>
//#include <libavformat/avformat.h>
import "C"

func GetFDFromFormatContext(fmtCtx *astiav.FormatContext) C.int {
	pb := fmtCtx.Pb()
	avioCtxR := reflect.ValueOf(unsafetools.FieldByName(pb, "c"))
	avioCtx := (*C.AVIOContext)(unsafe.Pointer(avioCtxR.Elem().Pointer()))
	urlCtx := (*C.URLContext)(avioCtx.opaque)
	srtCtx := (*C.SRTContext)(urlCtx.priv_data)
	return srtCtx.fd
}
