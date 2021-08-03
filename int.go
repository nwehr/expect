package expect

import "testing"

type intAssertion struct {
	t *testing.T
	n int
}

func (a intAssertion) ToEqual(n int) {
	if a.n != n {
		a.t.Fatalf("expected %d; got %d", n, a.n)
	}
}

func (a intAssertion) ToBeGreaterThan(n int) {
	if a.n <= n {
		a.t.Fatalf("expected %d to be greater than %d", a.n, n)
	}
}

func (a intAssertion) ToBeLessThan(n int) {
	if a.n >= n {
		a.t.Fatalf("expected %d to be less than %d", a.n, n)
	}
}
