// Copyright (C) 2015  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package endian_test

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/tsuna/endian"
)

func TestConversion(t *testing.T) {
	buf := make([]byte, 8)

	*(*uint16)(unsafe.Pointer(&buf[0])) = endian.HostToNetUint16(0xAABB)
	if !bytes.Equal(buf[:2], []byte("\xAA\xBB")) {
		t.Errorf("Invalid conversion yielded %x", buf[:2])
	}
	u16 := endian.NetToHostUint16(*(*uint16)(unsafe.Pointer(&buf[0])))
	if u16 != 0xAABB {
		t.Errorf("Invalid conversion yielded 0x%x", u16)
	}

	*(*uint32)(unsafe.Pointer(&buf[0])) = endian.HostToNetUint32(0xAABBCCDD)
	if !bytes.Equal(buf[:4], []byte("\xAA\xBB\xCC\xDD")) {
		t.Errorf("Invalid conversion yielded %x", buf[:4])
	}
	u32 := endian.NetToHostUint32(*(*uint32)(unsafe.Pointer(&buf[0])))
	if u32 != 0xAABBCCDD {
		t.Errorf("Invalid conversion yielded 0x%x", u32)
	}

	*(*uint64)(unsafe.Pointer(&buf[0])) = endian.HostToNetUint64(0xAABBCCDDEEFF4288)
	if !bytes.Equal(buf, []byte("\xAA\xBB\xCC\xDD\xEE\xFF\x42\x88")) {
		t.Errorf("Invalid conversion yielded %x", buf)
	}
	u64 := endian.NetToHostUint64(*(*uint64)(unsafe.Pointer(&buf[0])))
	if u64 != 0xAABBCCDDEEFF4288 {
		t.Errorf("Invalid conversion yielded 0x%x", u64)
	}
}
