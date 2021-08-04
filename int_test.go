package expect

import "testing"

func TestIntToEqual(t *testing.T) {
	Fatal(t).Int(1).ToEqual(1)
	Fatal(t).Int(2).ToEqual(2)
	Fatal(t).Int(3).ToEqual(3)
}

func TestIntToBeGreaterThan(t *testing.T) {
	Fatal(t).Int(10).ToBeGreaterThan(9)
	Fatal(t).Int(10).ToBeGreaterThan(0)
	Fatal(t).Int(10).ToBeGreaterThan(-3)
}

func TestIntToBeLessThan(t *testing.T) {
	Fatal(t).Int(8).ToBeLessThan(10)
	Fatal(t).Int(0).ToBeLessThan(10)
	Fatal(t).Int(-5).ToBeLessThan(10)
}
