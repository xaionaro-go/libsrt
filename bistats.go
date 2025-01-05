package libsrt

// #include <srt/srt.h>
import "C"

// See https://github.com/Haivision/srt/blob/8a89a3abbf4d3a2f7869d535349a474607ea0214/docs/API/statistics.md#srt-socket-statistics
func (s *Socket) Bistats(
	clear bool,
	instantaneous bool,
) (*TRACEBSTATS, error) {
	var stats TRACEBSTATS
	if ret := C.srt_bistats(s.C, &stats.C, boolToInt(clear), boolToInt(instantaneous)); ret != 0 {
		return nil, SomeError{}
	}

	return &stats, nil
}
