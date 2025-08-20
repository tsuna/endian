// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build 386 amd64,noasm amd64p32 arm arm64 ppc64le mipsle mips64le mips64p32le

package endian

// swapUint64 converts a uint16 to network byte order and back.
func swapUint64(n uint64) uint64 {
	return ((n & 0x00000000000000FF) << 56) |
		((n & 0x000000000000FF00) << 40) |
		((n & 0x0000000000FF0000) << 24) |
		((n & 0x00000000FF000000) << 8) |
		((n & 0x000000FF00000000) >> 8) |
		((n & 0x0000FF0000000000) >> 24) |
		((n & 0x00FF000000000000) >> 40) |
		((n & 0xFF00000000000000) >> 56)
}
