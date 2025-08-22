// Copyright (C) 2025  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build !noasm

#include "textflag.h"

#define SWAPUINT64 \
	MOVD n+0(FP), R0;   \
	REV R0, R1;         \
	MOVD R1, ret+8(FP); \
	RET

// func NetToHostUint64(n uint64) uint64
TEXT ·NetToHostUint64(SB), NOSPLIT, $0
	SWAPUINT64

// func HostToNetUint64(n uint64) uint64
TEXT ·HostToNetUint64(SB), NOSPLIT, $0
	SWAPUINT64
