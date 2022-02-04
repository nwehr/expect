package expect

import (
	"testing"
)

type assert struct {
	t        *testing.T
	severity severity
}

func (e assert) Int(n int) intAssertion {
	e.t.Helper()

	return intAssertion{
		n: n,
		t: e.t,
	}
}

func (e assert) UInt(n uint) uintAssertion {
	e.t.Helper()

	return uintAssertion{
		n: n,
		t: e.t,
	}
}

func (e assert) String(s string) stringAssertion {
	e.t.Helper()

	return stringAssertion{
		s: s,
		t: e.t,
	}
}

func (e assert) NoError(err error) {
	e.t.Helper()

	if err != nil {
		e.t.Errorf("unexpected error; %v", err)
	}
}

func (e assert) Error(err error) {
	e.t.Helper()

	if err == nil {
		e.t.Errorf("expected error")
	}
}
