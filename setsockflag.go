package libsrt

// #include <srt/srt.h>
import "C"
import "runtime"

type SockoptValue interface {
	Blob
}

// https://github.com/Haivision/srt/blob/master/docs/API/API-functions.md#srt_setsockflag
// int srt_setsockflag(SRTSOCKET u, SRT_SOCKOPT opt, const void* optval, int optlen);
func (s *Socket) Setsockflag(
	opt Sockopt,
	value SockoptValue,
) error {
	ptr := value.Pointer()
	rc := C.srt_setsockflag(s.C, C.SRT_SOCKOPT(opt), ptr, C.int(value.Size()))
	if rc != 0 {
		return SomeError{}
	}
	runtime.KeepAlive(ptr)
	return nil
}
