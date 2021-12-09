package expect

import (
	"testing"
)

func TestNoError(t *testing.T) {
	T(t).NoError(nil)
}
