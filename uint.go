package expect

type uintAssertion struct {
	n     uint
	failf func(string, ...interface{})
}

func (a uintAssertion) ToEqual(n uint) {
	if a.n != n {
		a.failf("expected %d; got %d", n, a.n)
	}
}

func (a uintAssertion) ToBeGreaterThan(n uint) {
	if a.n <= n {
		a.failf("expected %d to be greater than %d", a.n, n)
	}
}

func (a uintAssertion) ToBeLessThan(n uint) {
	if a.n >= n {
		a.failf("expected %d to be less than %d", a.n, n)
	}
}
