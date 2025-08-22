// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build 386 amd64,noasm amd64p32 arm arm64,noasm ppc64le mipsle mips64le mips64p32le

package endian

import _ "unsafe"

//go:linkname netToHostUint32 github.com/tsuna/endian.NetToHostUint32
func netToHostUint32(n uint32) uint32 {
	return swapUint32(n)
}

//go:linkname hostToNetUint32 github.com/tsuna/endian.HostToNetUint32
func hostToNetUint32(n uint32) uint32 {
	return swapUint32(n)
}

// swapUint32 converts a uint16 to network byte order and back.
func swapUint32(n uint32) uint32 {
	return (n&0x000000FF)<<24 | (n&0x0000FF00)<<8 |
		(n&0x00FF0000)>>8 | (n&0xFF000000)>>24
}
