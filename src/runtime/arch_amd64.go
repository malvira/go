// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	thechar        = '6'
	_BigEndian     = 0
	_CacheLineSize = 64
	_PhysPageSize  = 4096
	_PCQuantum     = 1
	_Int64Align    = 8
	hugePageSize   = 1 << 21
	minFrameSize   = 0
)

type uintreg uint64
type intptr int64 // TODO(rsc): remove
