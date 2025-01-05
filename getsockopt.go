package libsrt

// #include <srt/srt.h>
import "C"
import "runtime"

// https://github.com/Haivision/srt/blob/8a89a3abbf4d3a2f7869d535349a474607ea0214/docs/API/API-functions.md#srt_getsockopt
// int srt_getsockopt(SRTSOCKET u, int level /*ignored*/, SRT_SOCKOPT opt, void* optval, int* optlen);
// DEPRECATED: Use Getsockflag instead.
func (s *Socket) Getsockopt(
	level int,
	opt Sockopt,
	value SockoptValueWritable,
) error {
	ptr := value.Pointer()
	rc := C.srt_getsockopt(s.C, C.int(level), C.SRT_SOCKOPT(opt), ptr, (*C.int)(value.SizePointer()))
	if rc != 0 {
		return SomeError{}
	}
	runtime.KeepAlive(ptr)
	return nil
}
