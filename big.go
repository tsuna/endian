// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build armbe arm64be ppc64 mips mips64 mips64p32 ppc s390 s390x sparc sparc64

package endian

import _ "unsafe"

//go:linkname netToHostUint16 github.com/tsuna/endian.NetToHostUint16
func netToHostUint16(n uint16) uint16 {
	return n
}

//go:linkname hostToNetUint16 github.com/tsuna/endian.HostToNetUint16
func hostToNetUint16(n uint16) uint16 {
	return n
}

//go:linkname netToHostUint32 github.com/tsuna/endian.NetToHostUint32
func netToHostUint32(n uint32) uint32 {
	return n
}

//go:linkname hostToNetUint32 github.com/tsuna/endian.HostToNetUint32
func hostToNetUint32(n uint32) uint32 {
	return n
}

//go:linkname netToHostUint64 github.com/tsuna/endian.NetToHostUint64
func netToHostUint64(n uint64) uint64 {
	return n
}

//go:linkname hostToNetUint64 github.com/tsuna/endian.HostToNetUint64
func hostToNetUint64(n uint64) uint64 {
	return n
}
