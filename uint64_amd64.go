// Copyright (C) 2025  The endian Authors.  All rights reserved.
// This file is part of the Go endian library.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// +build !noasm

package endian

// implemented in uint64_amd64.s
func swapUint64(n uint64) uint64
