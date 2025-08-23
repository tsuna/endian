// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build 386 amd64 amd64p32 arm arm64 ppc64le mipsle mips64le mips64p32le

package endian

import _ "unsafe"

//go:linkname netToHostUint16 github.com/tsuna/endian.NetToHostUint16
func netToHostUint16(n uint16) uint16 {
	return swapUint16(n)
}

//go:linkname hostToNetUint16 github.com/tsuna/endian.HostToNetUint16
func hostToNetUint16(n uint16) uint16 {
	return swapUint16(n)
}

// swapUint16 converts a uint16 to network byte order and back.
func swapUint16(n uint16) uint16 {
	return (n&0x00FF)<<8 | (n&0xFF00)>>8
}
