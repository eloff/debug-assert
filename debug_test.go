// +build debug

package assert

import (
	"fmt"
	"testing"

	"github.com/eloff/assert"
)

func TestCompares(t *testing.T) {
	type compares func(x, y interface{})

	tests := []struct {
		op  string
		cmp compares
		// Truth table for
		// cmp(0, 0)
		// cmp(0, 1)
		// cmp(0, 2)
		// cmp(1, 0)
		// cmp(1, 1)
		// cmp(1, 2)
		// cmp(2, 0)
		// cmp(2, 1)
		// cmp(2, 2)
		expected        [9]bool
		invertedMessage bool
	}{
		{
			op:       ">",
			cmp:      Greater,
			expected: [9]bool{false, false, false, true, false, false, true, true, false},
		},
		{
			op:       "<",
			cmp:      Lesser,
			expected: [9]bool{false, true, true, false, false, true, false, false, false},
		},
		{
			op:              "==",
			cmp:             func(x, y interface{}) { Equal(x, y) },
			expected:        [9]bool{true, false, false, false, true, false, false, false, true},
			invertedMessage: true,
		},
		{
			op:              "!=",
			cmp:             func(x, y interface{}) { NotEqual(x, y) },
			expected:        [9]bool{false, true, true, true, false, true, true, true, false},
			invertedMessage: true,
		},
		{
			op:       ">=",
			cmp:      GreaterOrEqual,
			expected: [9]bool{true, false, false, true, true, false, true, true, true},
		},
		{
			op:       "<=",
			cmp:      LesserOrEqual,
			expected: [9]bool{true, true, true, false, true, true, false, false, true},
		},
	}

	for _, test := range tests {
		a := assert.New(t, test.op)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				f := func() { test.cmp(i, j) }
				var msg string
				if test.invertedMessage {
					msg = fmt.Sprintf("%d %s %d", j, test.op, i)
				} else {
					msg = fmt.Sprintf("%d %s %d", i, test.op, j)
				}
				t.Log(msg)
				if !test.expected[i*3+j] {
					a.Panics(f, msg)
				} else {
					f()
				}
			}
		}
	}
}
