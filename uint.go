package expect

import "testing"

type uintAssertion struct {
	n uint
	t *testing.T
}

func (a uintAssertion) ToEqual(n uint) {
	a.t.Helper()

	if a.n != n {
		a.t.Errorf("expected %d; got %d", n, a.n)
	}
}

func (a uintAssertion) ToBeGreaterThan(n uint) {
	a.t.Helper()

	if a.n <= n {
		a.t.Errorf("expected %d to be greater than %d", a.n, n)
	}
}

func (a uintAssertion) ToBeLessThan(n uint) {
	a.t.Helper()

	if a.n >= n {
		a.t.Errorf("expected %d to be less than %d", a.n, n)
	}
}
