package expect

import (
	"fmt"
	"testing"
)

type assert struct {
	t        *testing.T
	severity severity
}

func (e assert) Int(n int) intAssertion {
	return intAssertion{
		n: n,
		failf: func(format string, args ...interface{}) {
			fail(e.t, fmt.Sprintf(format, args...), e.severity)
		},
	}
}

func (e assert) UInt(n uint) uintAssertion {
	return uintAssertion{
		n: n,
		failf: func(format string, args ...interface{}) {
			fail(e.t, fmt.Sprintf(format, args...), e.severity)
		},
	}
}

func (e assert) String(s string) stringAssertion {
	return stringAssertion{t: e.t, s: s}
}

func (e assert) NoError(err error) {
	if err != nil {
		e.t.Errorf("unexpected error: %s", err)
	}
}

func (e assert) Error(err error) {
	if err == nil {
		e.t.Errorf("expected error; got nil")
	}
}
