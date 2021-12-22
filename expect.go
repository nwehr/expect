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
	return assert{
		t:        t,
		severity: _error,
	}
}

func fail(t *testing.T, msg string, sev severity) {
	if sev == _fatal {
		t.Fatal(msg)
	}

	if sev == _error {
		t.Error(msg)
	}
}
