package libsrt

// #include <srt/srt.h>
import "C"

// See https://github.com/Haivision/srt/blob/8a89a3abbf4d3a2f7869d535349a474607ea0214/docs/API/API-functions.md#srt_strerror
// const char* srt_strerror(int code, int errnoval);
func Strerror(code CodeMinor, errno Errno) string {
	s := C.srt_strerror(C.int(code), C.int(errno))
	return C.GoString(s)
}
