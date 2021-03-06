// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simple

import (
	"go/parser"

	"github.com/gwaylib/godep/gps"
)

var (
	_ = parser.ParseFile
	S = gps.Prepare
)
