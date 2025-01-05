package libsrt

// #include <srt/srt.h>
import "C"
import "fmt"

type SomeError struct{}

func (SomeError) Error() string {
	// https://github.com/Haivision/srt/blob/master/docs/API/API-functions.md#srt_getlasterror
	return "an error, use Getlasterror() to obtain the specific error"
}

type Error struct {
	Code  CodeMinor
	Errno Errno
}

func (err Error) Error() string {
	return fmt.Sprintf("code:%d, errno:%d: %s", err.Code, err.Errno, Strerror(err.Code, err.Errno))
}
