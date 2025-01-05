package sockopt

// #include <srt/srt.h>
import "C"
import (
	"fmt"
	"strings"
)

type Sockopt C.SRT_SOCKOPT

const (
	MSS                 = Sockopt(C.SRTO_MSS)
	SNDSYN              = Sockopt(C.SRTO_SNDSYN)
	RCVSYN              = Sockopt(C.SRTO_RCVSYN)
	ISN                 = Sockopt(C.SRTO_ISN)
	FC                  = Sockopt(C.SRTO_FC)
	SNDBUF              = Sockopt(C.SRTO_SNDBUF)
	RCVBUF              = Sockopt(C.SRTO_RCVBUF)
	LINGER              = Sockopt(C.SRTO_LINGER)
	UDP_SNDBUF          = Sockopt(C.SRTO_UDP_SNDBUF)
	UDP_RCVBUF          = Sockopt(C.SRTO_UDP_RCVBUF)
	RENDEZVOUS          = Sockopt(C.SRTO_RENDEZVOUS)
	SNDTIMEO            = Sockopt(C.SRTO_SNDTIMEO)
	RCVTIMEO            = Sockopt(C.SRTO_RCVTIMEO)
	REUSEADDR           = Sockopt(C.SRTO_REUSEADDR)
	MAXBW               = Sockopt(C.SRTO_MAXBW)
	STATE               = Sockopt(C.SRTO_STATE)
	EVENT               = Sockopt(C.SRTO_EVENT)
	SNDDATA             = Sockopt(C.SRTO_SNDDATA)
	RCVDATA             = Sockopt(C.SRTO_RCVDATA)
	SENDER              = Sockopt(C.SRTO_SENDER)
	TSBPDMODE           = Sockopt(C.SRTO_TSBPDMODE)
	LATENCY             = Sockopt(C.SRTO_LATENCY)
	INPUTBW             = Sockopt(C.SRTO_INPUTBW)
	OHEADBW             = Sockopt(C.SRTO_OHEADBW)
	PASSPHRASE          = Sockopt(C.SRTO_PASSPHRASE)
	PBKEYLEN            = Sockopt(C.SRTO_PBKEYLEN)
	KMSTATE             = Sockopt(C.SRTO_KMSTATE)
	IPTTL               = Sockopt(C.SRTO_IPTTL)
	IPTOS               = Sockopt(C.SRTO_IPTOS)
	TLPKTDROP           = Sockopt(C.SRTO_TLPKTDROP)
	SNDDROPDELAY        = Sockopt(C.SRTO_SNDDROPDELAY)
	NAKREPORT           = Sockopt(C.SRTO_NAKREPORT)
	VERSION             = Sockopt(C.SRTO_VERSION)
	PEERVERSION         = Sockopt(C.SRTO_PEERVERSION)
	CONNTIMEO           = Sockopt(C.SRTO_CONNTIMEO)
	DRIFTTRACER         = Sockopt(C.SRTO_DRIFTTRACER)
	MININPUTBW          = Sockopt(C.SRTO_MININPUTBW)
	SNDKMSTATE          = Sockopt(C.SRTO_SNDKMSTATE)
	RCVKMSTATE          = Sockopt(C.SRTO_RCVKMSTATE)
	LOSSMAXTTL          = Sockopt(C.SRTO_LOSSMAXTTL)
	RCVLATENCY          = Sockopt(C.SRTO_RCVLATENCY)
	PEERLATENCY         = Sockopt(C.SRTO_PEERLATENCY)
	MINVERSION          = Sockopt(C.SRTO_MINVERSION)
	STREAMID            = Sockopt(C.SRTO_STREAMID)
	CONGESTION          = Sockopt(C.SRTO_CONGESTION)
	MESSAGEAPI          = Sockopt(C.SRTO_MESSAGEAPI)
	PAYLOADSIZE         = Sockopt(C.SRTO_PAYLOADSIZE)
	TRANSTYPE           = Sockopt(C.SRTO_TRANSTYPE)
	KMREFRESHRATE       = Sockopt(C.SRTO_KMREFRESHRATE)
	KMPREANNOUNCE       = Sockopt(C.SRTO_KMPREANNOUNCE)
	ENFORCEDENCRYPTION  = Sockopt(C.SRTO_ENFORCEDENCRYPTION)
	IPV6ONLY            = Sockopt(C.SRTO_IPV6ONLY)
	PEERIDLETIMEO       = Sockopt(C.SRTO_PEERIDLETIMEO)
	BINDTODEVICE        = Sockopt(C.SRTO_BINDTODEVICE)
	GROUPCONNECT        = Sockopt(C.SRTO_GROUPCONNECT)
	GROUPMINSTABLETIMEO = Sockopt(C.SRTO_GROUPMINSTABLETIMEO)
	GROUPTYPE           = Sockopt(C.SRTO_GROUPTYPE)
	PACKETFILTER        = Sockopt(C.SRTO_PACKETFILTER)
	RETRANSMITALGO      = Sockopt(C.SRTO_RETRANSMITALGO)
	E_SIZE              = Sockopt(C.SRTO_E_SIZE)
)

func (opt Sockopt) String() string {
	switch opt {
	case MSS:
		return "MSS"
	case SNDSYN:
		return "SNDSYN"
	case RCVSYN:
		return "RCVSYN"
	case ISN:
		return "ISN"
	case FC:
		return "FC"
	case SNDBUF:
		return "SNDBUF"
	case RCVBUF:
		return "RCVBUF"
	case LINGER:
		return "LINGER"
	case UDP_SNDBUF:
		return "UDP_SNDBUF"
	case UDP_RCVBUF:
		return "UDP_RCVBUF"
	case RENDEZVOUS:
		return "RENDEZVOUS"
	case SNDTIMEO:
		return "SNDTIMEO"
	case RCVTIMEO:
		return "RCVTIMEO"
	case REUSEADDR:
		return "REUSEADDR"
	case MAXBW:
		return "MAXBW"
	case STATE:
		return "STATE"
	case EVENT:
		return "EVENT"
	case SNDDATA:
		return "SNDDATA"
	case RCVDATA:
		return "RCVDATA"
	case SENDER:
		return "SENDER"
	case TSBPDMODE:
		return "TSBPDMODE"
	case LATENCY:
		return "LATENCY"
	case INPUTBW:
		return "INPUTBW"
	case OHEADBW:
		return "OHEADBW"
	case PASSPHRASE:
		return "PASSPHRASE"
	case PBKEYLEN:
		return "PBKEYLEN"
	case KMSTATE:
		return "KMSTATE"
	case IPTTL:
		return "IPTTL"
	case IPTOS:
		return "IPTOS"
	case TLPKTDROP:
		return "TLPKTDROP"
	case SNDDROPDELAY:
		return "SNDDROPDELAY"
	case NAKREPORT:
		return "NAKREPORT"
	case VERSION:
		return "VERSION"
	case PEERVERSION:
		return "PEERVERSION"
	case CONNTIMEO:
		return "CONNTIMEO"
	case DRIFTTRACER:
		return "DRIFTTRACER"
	case MININPUTBW:
		return "MININPUTBW"
	case SNDKMSTATE:
		return "SNDKMSTATE"
	case RCVKMSTATE:
		return "RCVKMSTATE"
	case LOSSMAXTTL:
		return "LOSSMAXTTL"
	case RCVLATENCY:
		return "RCVLATENCY"
	case PEERLATENCY:
		return "PEERLATENCY"
	case MINVERSION:
		return "MINVERSION"
	case STREAMID:
		return "STREAMID"
	case CONGESTION:
		return "CONGESTION"
	case MESSAGEAPI:
		return "MESSAGEAPI"
	case PAYLOADSIZE:
		return "PAYLOADSIZE"
	case TRANSTYPE:
		return "TRANSTYPE"
	case KMREFRESHRATE:
		return "KMREFRESHRATE"
	case KMPREANNOUNCE:
		return "KMPREANNOUNCE"
	case ENFORCEDENCRYPTION:
		return "ENFORCEDENCRYPTION"
	case IPV6ONLY:
		return "IPV6ONLY"
	case PEERIDLETIMEO:
		return "PEERIDLETIMEO"
	case BINDTODEVICE:
		return "BINDTODEVICE"
	case GROUPCONNECT:
		return "GROUPCONNECT"
	case GROUPMINSTABLETIMEO:
		return "GROUPMINSTABLETIMEO"
	case GROUPTYPE:
		return "GROUPTYPE"
	case PACKETFILTER:
		return "PACKETFILTER"
	case RETRANSMITALGO:
		return "RETRANSMITALGO"
	case E_SIZE:
		return "E_SIZE"
	}
	return fmt.Sprintf("unexpected_sockopt_%d", opt)
}

func FromString(s string) (Sockopt, bool) {
	switch strings.ToUpper(s) {
	case MSS.String():
		return MSS, true
	case SNDSYN.String():
		return SNDSYN, true
	case RCVSYN.String():
		return RCVSYN, true
	case ISN.String():
		return ISN, true
	case FC.String():
		return FC, true
	case SNDBUF.String():
		return SNDBUF, true
	case RCVBUF.String():
		return RCVBUF, true
	case LINGER.String():
		return LINGER, true
	case UDP_SNDBUF.String():
		return UDP_SNDBUF, true
	case UDP_RCVBUF.String():
		return UDP_RCVBUF, true
	case RENDEZVOUS.String():
		return RENDEZVOUS, true
	case SNDTIMEO.String():
		return SNDTIMEO, true
	case RCVTIMEO.String():
		return RCVTIMEO, true
	case REUSEADDR.String():
		return REUSEADDR, true
	case MAXBW.String():
		return MAXBW, true
	case STATE.String():
		return STATE, true
	case EVENT.String():
		return EVENT, true
	case SNDDATA.String():
		return SNDDATA, true
	case RCVDATA.String():
		return RCVDATA, true
	case SENDER.String():
		return SENDER, true
	case TSBPDMODE.String():
		return TSBPDMODE, true
	case LATENCY.String():
		return LATENCY, true
	case INPUTBW.String():
		return INPUTBW, true
	case OHEADBW.String():
		return OHEADBW, true
	case PASSPHRASE.String():
		return PASSPHRASE, true
	case PBKEYLEN.String():
		return PBKEYLEN, true
	case KMSTATE.String():
		return KMSTATE, true
	case IPTTL.String():
		return IPTTL, true
	case IPTOS.String():
		return IPTOS, true
	case TLPKTDROP.String():
		return TLPKTDROP, true
	case SNDDROPDELAY.String():
		return SNDDROPDELAY, true
	case NAKREPORT.String():
		return NAKREPORT, true
	case VERSION.String():
		return VERSION, true
	case PEERVERSION.String():
		return PEERVERSION, true
	case CONNTIMEO.String():
		return CONNTIMEO, true
	case DRIFTTRACER.String():
		return DRIFTTRACER, true
	case MININPUTBW.String():
		return MININPUTBW, true
	case SNDKMSTATE.String():
		return SNDKMSTATE, true
	case RCVKMSTATE.String():
		return RCVKMSTATE, true
	case LOSSMAXTTL.String():
		return LOSSMAXTTL, true
	case RCVLATENCY.String():
		return RCVLATENCY, true
	case PEERLATENCY.String():
		return PEERLATENCY, true
	case MINVERSION.String():
		return MINVERSION, true
	case STREAMID.String():
		return STREAMID, true
	case CONGESTION.String():
		return CONGESTION, true
	case MESSAGEAPI.String():
		return MESSAGEAPI, true
	case PAYLOADSIZE.String():
		return PAYLOADSIZE, true
	case TRANSTYPE.String():
		return TRANSTYPE, true
	case KMREFRESHRATE.String():
		return KMREFRESHRATE, true
	case KMPREANNOUNCE.String():
		return KMPREANNOUNCE, true
	case ENFORCEDENCRYPTION.String():
		return ENFORCEDENCRYPTION, true
	case IPV6ONLY.String():
		return IPV6ONLY, true
	case PEERIDLETIMEO.String():
		return PEERIDLETIMEO, true
	case BINDTODEVICE.String():
		return BINDTODEVICE, true
	case GROUPCONNECT.String():
		return GROUPCONNECT, true
	case GROUPMINSTABLETIMEO.String():
		return GROUPMINSTABLETIMEO, true
	case GROUPTYPE.String():
		return GROUPTYPE, true
	case PACKETFILTER.String():
		return PACKETFILTER, true
	case RETRANSMITALGO.String():
		return RETRANSMITALGO, true
	case E_SIZE.String():
		return E_SIZE, true
	}
	return Sockopt(0), false
}

// DONOTUSE_IsIntOpt is supposed to return true if the option expects an integer value.
// However I had no time to actually check these options, so it works incorrectly, yet.
func (opt Sockopt) DONOTUSE_IsIntOpt() bool {
	switch opt {
	case MSS:
		return true
	case SNDSYN:
		return true
	case RCVSYN:
		return true
	case ISN:
		return true
	case FC:
		return true
	case SNDBUF:
		return true
	case RCVBUF:
		return true
	case LINGER:
		return true
	case UDP_SNDBUF:
		return true
	case UDP_RCVBUF:
		return true
	case RENDEZVOUS:
		return true
	case SNDTIMEO:
		return true
	case RCVTIMEO:
		return true
	case REUSEADDR:
		return true
	case MAXBW:
		return true
	case STATE:
		return true
	case EVENT:
		return true
	case SNDDATA:
		return true
	case RCVDATA:
		return true
	case SENDER:
		return true
	case TSBPDMODE:
		return true
	case LATENCY:
		return true
	case INPUTBW:
		return true
	case OHEADBW:
		return true
	case PASSPHRASE:
		return false
	case PBKEYLEN:
		return true
	case KMSTATE:
		return true
	case IPTTL:
		return true
	case IPTOS:
		return true
	case TLPKTDROP:
		return true
	case SNDDROPDELAY:
		return true
	case NAKREPORT:
		return true
	case VERSION:
		return true
	case PEERVERSION:
		return true
	case CONNTIMEO:
		return true
	case DRIFTTRACER:
		return true
	case MININPUTBW:
		return true
	case SNDKMSTATE:
		return true
	case RCVKMSTATE:
		return true
	case LOSSMAXTTL:
		return true
	case RCVLATENCY:
		return true
	case PEERLATENCY:
		return true
	case MINVERSION:
		return true
	case STREAMID:
		return true
	case CONGESTION:
		return true
	case MESSAGEAPI:
		return true
	case PAYLOADSIZE:
		return true
	case TRANSTYPE:
		return true
	case KMREFRESHRATE:
		return true
	case KMPREANNOUNCE:
		return true
	case ENFORCEDENCRYPTION:
		return true
	case IPV6ONLY:
		return true
	case PEERIDLETIMEO:
		return true
	case BINDTODEVICE:
		return true
	case GROUPCONNECT:
		return true
	case GROUPMINSTABLETIMEO:
		return true
	case GROUPTYPE:
		return true
	case PACKETFILTER:
		return true
	case RETRANSMITALGO:
		return true
	case E_SIZE:
		return true
	}
	return false
}
