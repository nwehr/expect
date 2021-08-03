package expect

import "testing"

type intAssertion struct {
	t *testing.T
	n int
}

func (a intAssertion) ToBe(n int) {
	if a.n != n {
		a.t.Fatalf("expected %d; got %d", n, a.n)
	}
}
