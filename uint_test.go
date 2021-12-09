package expect

import "testing"

func TestUIntToEqual(t *testing.T) {
	Fatal(t).UInt(1).ToEqual(1)
	Fatal(t).UInt(2).ToEqual(2)
	Fatal(t).UInt(3).ToEqual(3)
}

func TestUIntToBeGreaterThan(t *testing.T) {
	Fatal(t).UInt(10).ToBeGreaterThan(9)
	Fatal(t).UInt(10).ToBeGreaterThan(0)
	Fatal(t).UInt(10).ToBeGreaterThan(0)
}

func TestUIntToBeLessThan(t *testing.T) {
	Fatal(t).UInt(8).ToBeLessThan(10)
	Fatal(t).UInt(0).ToBeLessThan(10)
	Fatal(t).UInt(0).ToBeLessThan(10)
}
