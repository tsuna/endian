// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build armbe arm64be ppc64 mips mips64 mips64p32 ppc s390 s390x sparc sparc64

package endian

// SwapUint16 converts a uint16 to network byte order and back.
func SwapUint16(n uint16) uint16 {
	return n
}

// SwapUint32 converts a uint16 to network byte order and back.
func SwapUint32(n uint32) uint32 {
	return n
}

// SwapUint64 converts a uint16 to network byte order and back.
func SwapUint64(n uint64) uint64 {
	return n
}
