# About
[![License: CC0-1.0](https://img.shields.io/badge/License-CC0%201.0-lightgrey.svg)](http://creativecommons.org/publicdomain/zero/1.0/)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/xaionaro-go/libsrt)](https://pkg.go.dev/github.com/xaionaro-go/libsrt@v0.0.0-20250105190141-9742bb9d1089#pkg-index)

`libsrt` is a CGo-wrapper for the [Haivision SRT-library](https://github.com/Haivision/srt/blob/master/docs/API/API-functions.md).

Function are added on-demand basis, Pull Requests (e.g. to improve the coverage) are welcomed.

# Example

```go
package myfancypkg

import (
    "github.com/xaionaro-go/libsrt"
    "github.com/xaionaro-go/libsrt/sockopt"
    "github.com/xaionaro-go/libsrt/extras/xastiav"
)

func myFancyFunc() {
    [...]

    // extracting the file descriptor from a go-astiav's structure:

    sockS := xastiav.GetFDFromFormatContext(formatCtx)

    // manipulating the SRT socket

    s := libsrt.SocketFromC(sockC)
    err := s.Setsockflag(sockopt.LATENCY, libsrt.BlobInt(5000000)) // 5 seconds
    if err != nil {
        return fmt.Errorf("unable to set latency to 5 seconds: %w", libsrt.Getlasterror()) // yes, this is how libsrt works, unfortunately, you need to ask for an error code in a thread-unsafe manner.
    }

    [...]
}
```