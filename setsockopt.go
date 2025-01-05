package libsrt

// #include <srt/srt.h>
import "C"
import "runtime"

// https://github.com/Haivision/srt/blob/8a89a3abbf4d3a2f7869d535349a474607ea0214/docs/API/API-functions.md#srt_setsockopt
// int srt_setsockopt(SRTSOCKET u, int level /*ignored*/, SRT_SOCKOPT opt, const void* optval, int optlen);
// DEPRECATED: Use Setsockflag instead.
func (s *Socket) Setsockopt(
	level int,
	opt Sockopt,
	value SockoptValue,
) error {
	ptr := value.Pointer()
	rc := C.srt_setsockopt(s.C, C.int(level), C.SRT_SOCKOPT(opt), ptr, value.Size())
	if rc != 0 {
		return SomeError{}
	}
	runtime.KeepAlive(ptr)
	return nil
}
