package expect

import "testing"

type intAssertion struct {
	n int
	t *testing.T
}

func (a intAssertion) ToEqual(n int) {
	a.t.Helper()

	if a.n != n {
		a.t.Errorf("expected %d; got %d", n, a.n)
	}
}

func (a intAssertion) ToBeGreaterThan(n int) {
	a.t.Helper()

	if a.n <= n {
		a.t.Errorf("expected %d to be greater than %d", n, a.n)
	}
}

func (a intAssertion) ToBeLessThan(n int) {
	a.t.Helper()

	if a.n >= n {
		a.t.Errorf("expected %d to be less than %d", n, a.n)
	}
}
