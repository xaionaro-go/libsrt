package libsrt

// #include <srt/srt.h>
import "C"

type TRACEBSTATS struct {
	C C.SRT_TRACEBSTATS
}

func (s *TRACEBSTATS) Convert() Tracebstats {
	if s == nil {
		return Tracebstats{}
	}
	return Tracebstats{
		MsTimeStamp:             int64(s.C.msTimeStamp),
		PktSentTotal:            int64(s.C.pktSentTotal),
		PktRecvTotal:            int64(s.C.pktRecvTotal),
		PktSndLossTotal:         int(s.C.pktSndLossTotal),
		PktRcvLossTotal:         int(s.C.pktRcvLossTotal),
		PktRetransTotal:         int(s.C.pktRetransTotal),
		PktSentACKTotal:         int(s.C.pktSentACKTotal),
		PktRecvACKTotal:         int(s.C.pktRecvACKTotal),
		PktSentNAKTotal:         int(s.C.pktSentNAKTotal),
		PktRecvNAKTotal:         int(s.C.pktRecvNAKTotal),
		UsSndDurationTotal:      int64(s.C.usSndDurationTotal),
		PktSndDropTotal:         int(s.C.pktSndDropTotal),
		PktRcvDropTotal:         int(s.C.pktRcvDropTotal),
		PktRcvUndecryptTotal:    int(s.C.pktRcvUndecryptTotal),
		ByteSentTotal:           uint64(s.C.byteSentTotal),
		ByteRecvTotal:           uint64(s.C.byteRecvTotal),
		ByteRcvLossTotal:        uint64(s.C.byteRcvLossTotal),
		ByteRetransTotal:        uint64(s.C.byteRetransTotal),
		ByteSndDropTotal:        uint64(s.C.byteSndDropTotal),
		ByteRcvDropTotal:        uint64(s.C.byteRcvDropTotal),
		ByteRcvUndecryptTotal:   uint64(s.C.byteRcvUndecryptTotal),
		PktSent:                 int64(s.C.pktSent),
		PktRecv:                 int64(s.C.pktRecv),
		PktSndLoss:              int(s.C.pktSndLoss),
		PktRcvLoss:              int(s.C.pktRcvLoss),
		PktRetrans:              int(s.C.pktRetrans),
		PktRcvRetrans:           int(s.C.pktRcvRetrans),
		PktSentACK:              int(s.C.pktSentACK),
		PktRecvACK:              int(s.C.pktRecvACK),
		PktSentNAK:              int(s.C.pktSentNAK),
		PktRecvNAK:              int(s.C.pktRecvNAK),
		MbpsSendRate:            float64(s.C.mbpsSendRate),
		MbpsRecvRate:            float64(s.C.mbpsRecvRate),
		UsSndDuration:           int64(s.C.usSndDuration),
		PktReorderDistance:      int(s.C.pktReorderDistance),
		PktRcvAvgBelatedTime:    float64(s.C.pktRcvAvgBelatedTime),
		PktRcvBelated:           int64(s.C.pktRcvBelated),
		PktSndDrop:              int(s.C.pktSndDrop),
		PktRcvDrop:              int(s.C.pktRcvDrop),
		PktRcvUndecrypt:         int(s.C.pktRcvUndecrypt),
		ByteSent:                uint64(s.C.byteSent),
		ByteRecv:                uint64(s.C.byteRecv),
		ByteRcvLoss:             uint64(s.C.byteRcvLoss),
		ByteRetrans:             uint64(s.C.byteRetrans),
		ByteSndDrop:             uint64(s.C.byteSndDrop),
		ByteRcvDrop:             uint64(s.C.byteRcvDrop),
		ByteRcvUndecrypt:        uint64(s.C.byteRcvUndecrypt),
		UsPktSndPeriod:          float64(s.C.usPktSndPeriod),
		PktFlowWindow:           int(s.C.pktFlowWindow),
		PktCongestionWindow:     int(s.C.pktCongestionWindow),
		PktFlightSize:           int(s.C.pktFlightSize),
		MsRTT:                   float64(s.C.msRTT),
		MbpsBandwidth:           float64(s.C.mbpsBandwidth),
		ByteAvailSndBuf:         int(s.C.byteAvailSndBuf),
		ByteAvailRcvBuf:         int(s.C.byteAvailRcvBuf),
		MbpsMaxBW:               float64(s.C.mbpsMaxBW),
		ByteMSS:                 int(s.C.byteMSS),
		PktSndBuf:               int(s.C.pktSndBuf),
		ByteSndBuf:              int(s.C.byteSndBuf),
		MsSndBuf:                int(s.C.msSndBuf),
		MsSndTsbPdDelay:         int(s.C.msSndTsbPdDelay),
		PktRcvBuf:               int(s.C.pktRcvBuf),
		ByteRcvBuf:              int(s.C.byteRcvBuf),
		MsRcvBuf:                int(s.C.msRcvBuf),
		MsRcvTsbPdDelay:         int(s.C.msRcvTsbPdDelay),
		PktSndFilterExtraTotal:  int(s.C.pktSndFilterExtraTotal),
		PktRcvFilterExtraTotal:  int(s.C.pktRcvFilterExtraTotal),
		PktRcvFilterSupplyTotal: int(s.C.pktRcvFilterSupplyTotal),
		PktRcvFilterLossTotal:   int(s.C.pktRcvFilterLossTotal),
		PktSndFilterExtra:       int(s.C.pktSndFilterExtra),
		PktRcvFilterExtra:       int(s.C.pktRcvFilterExtra),
		PktRcvFilterSupply:      int(s.C.pktRcvFilterSupply),
		PktRcvFilterLoss:        int(s.C.pktRcvFilterLoss),
		PktReorderTolerance:     int(s.C.pktReorderTolerance),
		PktSentUniqueTotal:      int64(s.C.pktSentUniqueTotal),
		PktRecvUniqueTotal:      int64(s.C.pktRecvUniqueTotal),
		ByteSentUniqueTotal:     uint64(s.C.byteSentUniqueTotal),
		ByteRecvUniqueTotal:     uint64(s.C.byteRecvUniqueTotal),
		PktSentUnique:           int64(s.C.pktSentUnique),
		PktRecvUnique:           int64(s.C.pktRecvUnique),
		ByteSentUnique:          uint64(s.C.byteSentUnique),
		ByteRecvUnique:          uint64(s.C.byteRecvUnique),
	}
}

type Tracebstats struct {
	MsTimeStamp             int64
	PktSentTotal            int64
	PktRecvTotal            int64
	PktSndLossTotal         int
	PktRcvLossTotal         int
	PktRetransTotal         int
	PktSentACKTotal         int
	PktRecvACKTotal         int
	PktSentNAKTotal         int
	PktRecvNAKTotal         int
	UsSndDurationTotal      int64
	PktSndDropTotal         int
	PktRcvDropTotal         int
	PktRcvUndecryptTotal    int
	ByteSentTotal           uint64
	ByteRecvTotal           uint64
	ByteRcvLossTotal        uint64
	ByteRetransTotal        uint64
	ByteSndDropTotal        uint64
	ByteRcvDropTotal        uint64
	ByteRcvUndecryptTotal   uint64
	PktSent                 int64
	PktRecv                 int64
	PktSndLoss              int
	PktRcvLoss              int
	PktRetrans              int
	PktRcvRetrans           int
	PktSentACK              int
	PktRecvACK              int
	PktSentNAK              int
	PktRecvNAK              int
	MbpsSendRate            float64
	MbpsRecvRate            float64
	UsSndDuration           int64
	PktReorderDistance      int
	PktRcvAvgBelatedTime    float64
	PktRcvBelated           int64
	PktSndDrop              int
	PktRcvDrop              int
	PktRcvUndecrypt         int
	ByteSent                uint64
	ByteRecv                uint64
	ByteRcvLoss             uint64
	ByteRetrans             uint64
	ByteSndDrop             uint64
	ByteRcvDrop             uint64
	ByteRcvUndecrypt        uint64
	UsPktSndPeriod          float64
	PktFlowWindow           int
	PktCongestionWindow     int
	PktFlightSize           int
	MsRTT                   float64
	MbpsBandwidth           float64
	ByteAvailSndBuf         int
	ByteAvailRcvBuf         int
	MbpsMaxBW               float64
	ByteMSS                 int
	PktSndBuf               int
	ByteSndBuf              int
	MsSndBuf                int
	MsSndTsbPdDelay         int
	PktRcvBuf               int
	ByteRcvBuf              int
	MsRcvBuf                int
	MsRcvTsbPdDelay         int
	PktSndFilterExtraTotal  int
	PktRcvFilterExtraTotal  int
	PktRcvFilterSupplyTotal int
	PktRcvFilterLossTotal   int
	PktSndFilterExtra       int
	PktRcvFilterExtra       int
	PktRcvFilterSupply      int
	PktRcvFilterLoss        int
	PktReorderTolerance     int
	PktSentUniqueTotal      int64
	PktRecvUniqueTotal      int64
	ByteSentUniqueTotal     uint64
	ByteRecvUniqueTotal     uint64
	PktSentUnique           int64
	PktRecvUnique           int64
	ByteSentUnique          uint64
	ByteRecvUnique          uint64
}

// See https://github.com/Haivision/srt/blob/8a89a3abbf4d3a2f7869d535349a474607ea0214/docs/API/statistics.md#srt-socket-statistics
func (s *Socket) Bstats(clear bool) (*TRACEBSTATS, error) {
	var stats TRACEBSTATS
	if ret := C.srt_bstats(s.C, &stats.C, boolToInt(clear)); ret != 0 {
		return nil, SomeError{}
	}

	return &stats, nil
}
