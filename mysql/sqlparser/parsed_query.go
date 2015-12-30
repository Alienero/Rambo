// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"reflect"
)

type bindLocation struct {
	offset, length int
}

// ParsedQuery represents a parsed query where
// bind locations are precompued for fast substitutions.
type ParsedQuery struct {
	Query         string
	bindLocations []bindLocation
}

func IsNodeHasValue(node SQLNode) bool {
	t := reflect.TypeOf(node)
	v := reflect.ValueOf(node)
	switch t.Kind() {
	case reflect.Slice:
		if v.IsNil() || v.Len() < 1 {
			return false
		}
		return true
	case reflect.Struct:
		return true
	case reflect.Ptr:
		if v.IsNil() {
			return false
		}
		return true
	case reflect.String:
		if v.Len() < 1 {
			return false
		}
		return true
	default:
		return false
	}
}
