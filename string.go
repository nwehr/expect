package expect

import (
	"regexp"
	"testing"
)

type stringAssertion struct {
	t *testing.T
	s string
}

func (a stringAssertion) ToEqual(s string) {
	a.t.Helper()

	if a.s != s {
		a.t.Fatalf("expected '%s'; got '%s'", s, a.s)
	}
}

func (a stringAssertion) ToMatchPattern(pattern string) {
	a.t.Helper()

	matched, err := regexp.MatchString(pattern, a.s)
	if err != nil {
		a.t.Fatalf("unexpected error: %s", err.Error())
	}

	if !matched {
		a.t.Fatalf("expected '%s' to match '%s'", pattern, a.s)
	}
}
