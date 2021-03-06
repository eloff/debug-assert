// +build assert

package assert

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Assertf(msg string, args ...interface{}) {
	panic(fmt.Sprintf("Assert: "+msg, args...))
}

func Equal(expected, actual interface{}, opts ...cmp.Option) {
	opts = append(opts, cmpopts.EquateEmpty())
	if !cmp.Equal(expected, actual, opts...) {
		diff := fmt.Sprintf("expected %s\ngot %s\n", spew.Sdump(expected), spew.Sdump(actual))
		if len(diff) > 200 {
			// Try to get a shorter diff that will be more readable
			diff = fmt.Sprintf("expected %v\ngot %v\n", expected, actual)
			if len(diff) > 200 {
				diff = cmp.Diff(actual, expected, opts...)
			}
		}
		if diff != "" {
			Assertf("expected %v == %v: (-got +want)\n%s", actual, expected, diff)
		}
	}
}

func NotEqual(expected, actual interface{}, opts ...cmp.Option) {
	opts = append(opts, cmpopts.EquateEmpty())
	if cmp.Equal(expected, actual, opts...) {
		diff := spew.Sdump(expected)
		if len(diff) > 200 {
			// Try to get a shorter diff that will be more readable
			diff = fmt.Sprintf("%v", expected)
		}
		Assertf("actual equals expected:\n%s", diff)
	}
}

func NotNil(val interface{}) {
	if val == nil {
		Assertf("expected not nil")
	}
}

func True(val bool) {
	if !val {
		Assertf("expected true")
	}
}

func False(val bool) {
	if val {
		Assertf("expected false")
	}
}

func Greater(x, y interface{}) {
	if !less(y, x) {
		Assertf("expected %v > %v", x, y)
	}
}

func GreaterOrEqual(x, y interface{}) {
	if less(x, y) { //
		Assertf("expected %v >= %v", x, y)
	}
}

func Lesser(x, y interface{}) {
	if !less(x, y) {
		Assertf("expected %v < %v", x, y)
	}
}

func LesserOrEqual(x, y interface{}) {
	if less(y, x) { // x > y
		Assertf("expected %v <= %v", x, y)
	}
}

func TypesEqual(x, y interface{}) {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		Assertf("argument type mismatch: %T != %T", x, y)
	}
}

func less(x, y interface{}) bool {
	TypesEqual(x, y)
	switch val := x.(type) {
	case int:
		return val < y.(int)
	case uint:
		return val < y.(uint)
	case int32:
		return val < y.(int32)
	case uint32:
		return val < y.(uint32)
	case int64:
		return val < y.(int64)
	case uint64:
		return val < y.(uint64)
	case int16:
		return val < y.(int16)
	case uint16:
		return val < y.(uint16)
	case int8:
		return val < y.(int8)
	case uint8:
		return val < y.(uint8)
	case string:
		return val < y.(string)
	case float32:
		return val < y.(float32)
	case float64:
		return val < y.(float64)
	default:
		panic(fmt.Sprintf("type %T can't be compared with less than", val))
	}
}
