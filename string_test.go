package expect

import "testing"

func TestStringToEqual(t *testing.T) {
	T(t).String("abc123").ToEqual("abc123")
	T(t).String("123abc").ToEqual("123abc")
}

func TestStringMatchesPattern(t *testing.T) {
	T(t).String("abc123").ToMatchPattern("abc[0-9]+")
	T(t).String("abc123").ToMatchPattern("[a-z]+123")
}
