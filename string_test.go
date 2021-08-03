package expect

import "testing"

func TestStringToEqual(t *testing.T) {
	Expect(t).String("abc123").ToEqual("abc123")
	Expect(t).String("123abc").ToEqual("123abc")
}

func TestStringMatchesPattern(t *testing.T) {
	Expect(t).String("abc123").ToMatchPattern("abc[0-9]+")
	Expect(t).String("abc123").ToMatchPattern("[a-z]+123")
}
