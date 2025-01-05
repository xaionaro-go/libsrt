package libsrt

import (
	"sync"

	"github.com/xaionaro-go/libsrt"
)

var LibSrtLocker sync.Mutex

type Socket struct {
	Socket *libsrt.Socket
}

func SocketFromC(fd int32) *Socket {
	return WrapSocket(libsrt.SocketFromC(fd))
}

func WrapSocket(s *libsrt.Socket) *Socket {
	return &Socket{
		Socket: s,
	}
}

func (s *Socket) Bstats(
	clear bool,
) (*libsrt.TRACEBSTATS, error) {
	LibSrtLocker.Lock()
	defer LibSrtLocker.Unlock()

	result, err := s.Socket.Bstats(clear)
	if err != nil {
		err = libsrt.Getlasterror()
	}
	return result, err
}

func (s *Socket) Bistats(
	clear bool,
	instantaneous bool,
) (*libsrt.TRACEBSTATS, error) {
	LibSrtLocker.Lock()
	defer LibSrtLocker.Unlock()

	result, err := s.Socket.Bistats(clear, instantaneous)
	if err != nil {
		err = libsrt.Getlasterror()
	}
	return result, err
}

func (s *Socket) Setsockflag(
	opt libsrt.Sockopt,
	value libsrt.SockoptValue,
) error {
	LibSrtLocker.Lock()
	defer LibSrtLocker.Unlock()

	err := s.Socket.Setsockflag(opt, value)
	if err != nil {
		err = libsrt.Getlasterror()
	}
	return err
}

func (s *Socket) Setsockopt(
	limit int,
	opt libsrt.Sockopt,
	value libsrt.SockoptValue,
) error {
	LibSrtLocker.Lock()
	defer LibSrtLocker.Unlock()

	err := s.Socket.Setsockopt(limit, opt, value)
	if err != nil {
		err = libsrt.Getlasterror()
	}
	return err
}
