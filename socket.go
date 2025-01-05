package libsrt

// #include <srt/srt.h>
import "C"

type Socket struct {
	C C.SRTSOCKET
}

// SocketFromC wraps a file description of a socket into `Socket`.
//
// This could be useful for example if you work with libav (e.g. go-astiav) and you want
// to get access to an already existing socket.
func SocketFromC(fd int32) *Socket {
	return &Socket{
		C: C.SRTSOCKET(fd),
	}
}
