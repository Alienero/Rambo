// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

type bindLocation struct {
	offset, length int
}

// ParsedQuery represents a parsed query where
// bind locations are precompued for fast substitutions.
type ParsedQuery struct {
	Query         string
	bindLocations []bindLocation
}
