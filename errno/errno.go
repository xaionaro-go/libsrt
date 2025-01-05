package errno

// #include <srt/srt.h>
import "C"
import "fmt"

type Errno C.SRT_ERRNO

const (
	SUCCESS = Errno(C.SRT_SUCCESS)

	UNKNOWN        = Errno(C.SRT_EUNKNOWN)
	CONNSETUP      = Errno(C.SRT_ECONNSETUP)
	NOSERVER       = Errno(C.SRT_ENOSERVER)
	CONNREJ        = Errno(C.SRT_ECONNREJ)
	SOCKFAIL       = Errno(C.SRT_ESOCKFAIL)
	SECFAIL        = Errno(C.SRT_ESECFAIL)
	SCLOSED        = Errno(C.SRT_ESCLOSED)
	CONNFAIL       = Errno(C.SRT_ECONNFAIL)
	CONNLOST       = Errno(C.SRT_ECONNLOST)
	NOCONN         = Errno(C.SRT_ENOCONN)
	RESOURCE       = Errno(C.SRT_ERESOURCE)
	THREAD         = Errno(C.SRT_ETHREAD)
	NOBUF          = Errno(C.SRT_ENOBUF)
	SYSOBJ         = Errno(C.SRT_ESYSOBJ)
	FILE           = Errno(C.SRT_EFILE)
	INVRDOFF       = Errno(C.SRT_EINVRDOFF)
	RDPERM         = Errno(C.SRT_ERDPERM)
	INVWROFF       = Errno(C.SRT_EINVWROFF)
	WRPERM         = Errno(C.SRT_EWRPERM)
	INVOP          = Errno(C.SRT_EINVOP)
	BOUNDSOCK      = Errno(C.SRT_EBOUNDSOCK)
	CONNSOCK       = Errno(C.SRT_ECONNSOCK)
	INVPARAM       = Errno(C.SRT_EINVPARAM)
	INVSOCK        = Errno(C.SRT_EINVSOCK)
	UNBOUNDSOCK    = Errno(C.SRT_EUNBOUNDSOCK)
	NOLISTEN       = Errno(C.SRT_ENOLISTEN)
	RDVNOSERV      = Errno(C.SRT_ERDVNOSERV)
	RDVUNBOUND     = Errno(C.SRT_ERDVUNBOUND)
	INVALMSGAPI    = Errno(C.SRT_EINVALMSGAPI)
	INVALBUFFERAPI = Errno(C.SRT_EINVALBUFFERAPI)
	DUPLISTEN      = Errno(C.SRT_EDUPLISTEN)
	LARGEMSG       = Errno(C.SRT_ELARGEMSG)
	INVPOLLID      = Errno(C.SRT_EINVPOLLID)
	POLLEMPTY      = Errno(C.SRT_EPOLLEMPTY)
	BINDCONFLICT   = Errno(C.SRT_EBINDCONFLICT)
	ASYNCFAIL      = Errno(C.SRT_EASYNCFAIL)
	ASYNCSND       = Errno(C.SRT_EASYNCSND)
	ASYNCRCV       = Errno(C.SRT_EASYNCRCV)
	TIMEOUT        = Errno(C.SRT_ETIMEOUT)
	CONGEST        = Errno(C.SRT_ECONGEST)
	PEERERR        = Errno(C.SRT_EPEERERR)
)

func (e Errno) Error() string {
	return fmt.Sprintf("return code %d", int(e))
}
