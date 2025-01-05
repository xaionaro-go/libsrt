package threadsafe

import (
	"sync"

	"github.com/xaionaro-go/libsrt"
)

var LibSRTLocker sync.Mutex

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
	LibSRTLocker.Lock()
	defer LibSRTLocker.Unlock()

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
	LibSRTLocker.Lock()
	defer LibSRTLocker.Unlock()

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
	LibSRTLocker.Lock()
	defer LibSRTLocker.Unlock()

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
	LibSRTLocker.Lock()
	defer LibSRTLocker.Unlock()

	err := s.Socket.Setsockopt(limit, opt, value)
	if err != nil {
		err = libsrt.Getlasterror()
	}
	return err
}

func (s *Socket) Getsockopt(
	limit int,
	opt libsrt.Sockopt,
	value libsrt.SockoptValueWritable,
) error {
	LibSRTLocker.Lock()
	defer LibSRTLocker.Unlock()

	err := s.Socket.Getsockopt(limit, opt, value)
	if err != nil {
		err = libsrt.Getlasterror()
	}
	return err
}

func (s *Socket) Getsockflag(
	opt libsrt.Sockopt,
	value libsrt.SockoptValueWritable,
) error {
	LibSRTLocker.Lock()
	defer LibSRTLocker.Unlock()

	err := s.Socket.Getsockflag(opt, value)
	if err != nil {
		err = libsrt.Getlasterror()
	}
	return err
}
