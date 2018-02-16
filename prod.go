// +build !debug

package assert

import "github.com/google/go-cmp/cmp"

func Assertf(msg string, args ...interface{})                   {}
func Equal(expected, actual interface{}, opts ...cmp.Option)    {}
func NotEqual(expected, actual interface{}, opts ...cmp.Option) {}
func True(val bool)                                             {}
func False(val bool)                                            {}
func Greater(expected, actual interface{})                      {}
func GreaterOrEqual(expected, actual interface{})               {}
func Lesser(expected, actual interface{})                       {}
func LessserOrEqual(expected, actual interface{})               {}
func NotNil(val interface{})                                    {}
