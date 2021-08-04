package expect

type intAssertion struct {
	n     int
	failf func(string, ...interface{})
}

func (a intAssertion) ToEqual(n int) {
	if a.n != n {
		a.failf("expected %d; got %d", n, a.n)
	}
}

func (a intAssertion) ToBeGreaterThan(n int) {
	if a.n <= n {
		a.failf("expected %d to be greater than %d", a.n, n)
	}
}

func (a intAssertion) ToBeLessThan(n int) {
	if a.n >= n {
		a.failf("expected %d to be less than %d", a.n, n)
	}
}
