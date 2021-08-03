package expect

import "testing"

func TestIntToEqual(t *testing.T) {
	Expect(t).Int(1).ToEqual(1)
	Expect(t).Int(2).ToEqual(2)
	Expect(t).Int(3).ToEqual(3)
}

func TestIntToBeGreaterThan(t *testing.T) {
	Expect(t).Int(10).ToBeGreaterThan(9)
	Expect(t).Int(10).ToBeGreaterThan(0)
	Expect(t).Int(10).ToBeGreaterThan(-3)
}

func TestIntToBeLessThan(t *testing.T) {
	Expect(t).Int(8).ToBeLessThan(10)
	Expect(t).Int(0).ToBeLessThan(10)
	Expect(t).Int(-5).ToBeLessThan(10)
}
