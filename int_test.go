package expect

import "testing"

func TestIntEqualTo(t *testing.T) {
	Expect(t).Int(1).EqualTo(1)
	Expect(t).Int(2).EqualTo(2)
	Expect(t).Int(3).EqualTo(3)
}

func TestIntGreaterThan(t *testing.T) {
	Expect(t).Int(10).GreaterThan(9)
	Expect(t).Int(10).GreaterThan(0)
	Expect(t).Int(10).GreaterThan(-3)
}

func TestIntLessThan(t *testing.T) {
	Expect(t).Int(8).LessThan(10)
	Expect(t).Int(0).LessThan(10)
	Expect(t).Int(-5).LessThan(10)
}
