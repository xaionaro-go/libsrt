package libsrt

// #include <srt/srt.h>
import "C"
import "runtime"

type SockoptValueWritable interface {
	BlobWritable
}

// https://github.com/Haivision/srt/blob/8a89a3abbf4d3a2f7869d535349a474607ea0214/docs/API/API-functions.md#srt_getsockflag
// int srt_getsockflag(SRTSOCKET u, SRT_SOCKOPT opt, void* optval, int* optlen);
func (s *Socket) Getsockflag(
	opt Sockopt,
	value SockoptValueWritable,
) error {
	ptr := value.Pointer()
	rc := C.srt_getsockflag(s.C, C.SRT_SOCKOPT(opt), ptr, value.SizePointer())
	if rc != 0 {
		return SomeError{}
	}
	runtime.KeepAlive(ptr)
	return nil
}
