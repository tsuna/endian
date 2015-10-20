// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build armbe arm64be ppc64 mips mips64 mips64p32 ppc s390 s390x sparc sparc64

package endian

// swapUint16 converts a uint16 to network byte order and back.
func swapUint16(n uint16) uint16 {
	return n
}

// swapUint32 converts a uint16 to network byte order and back.
func swapUint32(n uint32) uint32 {
	return n
}

// swapUint64 converts a uint16 to network byte order and back.
func swapUint64(n uint64) uint64 {
	return n
}
