package expect

import "testing"

type expect struct {
	t *testing.T
}

func Expect(t *testing.T) expect {
	return expect{t: t}
}

func (e expect) Int(n int) intAssertion {
	return intAssertion{t: e.t, n: n}
}

func (e expect) String(s string) stringAssertion {
	return stringAssertion{t: e.t, s: s}
}
