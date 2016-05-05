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
	if node == nil {
		return false
	}
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
