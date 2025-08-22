// Copyright (C) 2025  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build !noasm

#include "textflag.h"

#define SWAPUINT32 \
	MOVW n+0(FP), R0;   \
	REV32 R0, R1;       \
	MOVW R1, ret+8(FP); \
	RET

// func NetToHostUint32(n uint32) uint32
TEXT ·NetToHostUint32(SB), NOSPLIT, $0
	SWAPUINT32

// func HostToNetUint32(n uint32) uint32
TEXT ·HostToNetUint32(SB), NOSPLIT, $0
	SWAPUINT32
