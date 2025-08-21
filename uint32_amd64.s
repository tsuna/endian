// Copyright (C) 2025  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build !noasm

#include "textflag.h"

// func NetToHostUint32(n uint32) uint32
TEXT ·NetToHostUint32(SB), NOSPLIT, $0
	MOVL n+0(FP), AX
	BSWAPL AX
	MOVL  AX, ret+8(FP)
	RET

// func HostToNetUint32(n uint32) uint32
TEXT ·HostToNetUint32(SB), NOSPLIT, $0
	MOVL n+0(FP), AX
	BSWAPL AX
	MOVL  AX, ret+8(FP)
	RET
