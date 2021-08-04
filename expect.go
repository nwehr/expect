package expect

import (
	"testing"
)

type severity int

const (
	_error severity = iota
	_fatal
)

func Error(t *testing.T) assert {
	return assert{
		t:        t,
		severity: _error,
	}
}

func Fatal(t *testing.T) assert {
	return assert{
		t:        t,
		severity: _fatal,
	}
}
