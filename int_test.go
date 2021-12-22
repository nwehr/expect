package expect

import "testing"

func TestIntToEqual(t *testing.T) {
	T(t).Int(1).ToEqual(1)
	T(t).Int(2).ToEqual(2)
	T(t).Int(3).ToEqual(3)
}

func TestIntToBeGreaterThan(t *testing.T) {
	T(t).Int(10).ToBeGreaterThan(9)
	T(t).Int(10).ToBeGreaterThan(0)
	T(t).Int(10).ToBeGreaterThan(-3)
}

func TestIntToBeLessThan(t *testing.T) {
	T(t).Int(8).ToBeLessThan(10)
	T(t).Int(0).ToBeLessThan(10)
	T(t).Int(-5).ToBeLessThan(10)
}
