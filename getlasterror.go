package libsrt

// #include <srt/srt.h>
import "C"

// https://github.com/Haivision/srt/blob/8a89a3abbf4d3a2f7869d535349a474607ea0214/docs/API/API-functions.md#srt_getlasterror
// int srt_getlasterror(int* errno_loc);
func Getlasterror() Error {
	var errnoC C.int
	code := C.srt_getlasterror(&errnoC)
	return Error{
		Code:  Code(code),
		Errno: Errno(errnoC),
	}
}
