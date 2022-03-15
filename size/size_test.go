package size

import (
	"fmt"
	"runtime"
	"testing"
)

type sizeTest struct {
	in  string
	out uint64
	err bool
}

var sizesTests = []sizeTest{
	{"", 0, false},
	{"10", 10, false},
	{"4k", 4096, false},
	{"10M", 10 * 1048576, false},
	{"80G", 80 * _GB, false},
	{"10T", 10 * _TB, false},

	{"4x", 0, true},
	{"boo", 0, true},
}

// make an assert() function for use in environment 't' and return it
func newAsserter(t *testing.T) func(cond bool, msg string, args ...interface{}) {
	return func(cond bool, msg string, args ...interface{}) {
		if cond {
			return
		}

		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "???"
			line = 0
		}

		s := fmt.Sprintf(msg, args...)
		t.Fatalf("%s: %d: Assertion failed: %s\n", file, line, s)
	}
}

func TestSize(t *testing.T) {
	assert := newAsserter(t)

	for i, t := range sizesTests {
		v, err := ParseSize(t.in)
		if t.err {
			assert(err != nil, "%2d: %s: expected to fail", i, t.in)
		} else {
			assert(err == nil, "%2d: %s: unexpected err: %s", i, t.in, err)
			assert(t.out == v, "%2d: %s: exp %v, saw %v", i, t.in, t.out, v)
		}
	}
}
