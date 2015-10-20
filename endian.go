// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// Package endian provides helper functions to deal with network byte order.
package endian

// NetToHostUint16 swaps the bytes of the given value (if necessary) to turn
// them from network byte order to host byte order.
func NetToHostUint16(n uint16) uint16 {
	return swapUint16(n)
}

// HostToNetUint16 swaps the bytes of the given value (if necessary) to turn
// them from network byte order to host byte order.
func HostToNetUint16(n uint16) uint16 {
	return swapUint16(n)
}

// NetToHostUint32 swaps the bytes of the given value (if necessary) to turn
// them from network byte order to host byte order.
func NetToHostUint32(n uint32) uint32 {
	return swapUint32(n)
}

// HostToNetUint32 swaps the bytes of the given value (if necessary) to turn
// them from network byte order to host byte order.
func HostToNetUint32(n uint32) uint32 {
	return swapUint32(n)
}

// NetToHostUint64 swaps the bytes of the given value (if necessary) to turn
// them from network byte order to host byte order.
func NetToHostUint64(n uint64) uint64 {
	return swapUint64(n)
}

// HostToNetUint64 swaps the bytes of the given value (if necessary) to turn
// them from network byte order to host byte order.
func HostToNetUint64(n uint64) uint64 {
	return swapUint64(n)
}
