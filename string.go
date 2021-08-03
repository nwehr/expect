package expect

import "testing"

type stringAssertion struct {
	t *testing.T
	s string
}

func (a stringAssertion) ToBe(s string) {
	if a.s != s {
		a.t.Fatalf("expected '%s'; got '%s'", s, a.s)
	}
}
