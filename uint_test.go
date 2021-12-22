package expect

import "testing"

func TestUIntToEqual(t *testing.T) {
	T(t).UInt(1).ToEqual(1)
	T(t).UInt(2).ToEqual(2)
	T(t).UInt(3).ToEqual(3)
}

func TestUIntToBeGreaterThan(t *testing.T) {
	T(t).UInt(10).ToBeGreaterThan(9)
	T(t).UInt(10).ToBeGreaterThan(0)
	T(t).UInt(10).ToBeGreaterThan(0)
}

func TestUIntToBeLessThan(t *testing.T) {
	T(t).UInt(8).ToBeLessThan(10)
	T(t).UInt(0).ToBeLessThan(10)
	T(t).UInt(0).ToBeLessThan(10)
}
