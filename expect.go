package expect

import (
	"testing"
)

type severity int

const (
	_error severity = iota
	_fatal
)

func T(t *testing.T) assert {
	t.Helper()

	return assert{
		t:        t,
		severity: _error,
	}
}
