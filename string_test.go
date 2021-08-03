package expect

import "testing"

func TestStringEqualTo(t *testing.T) {
	Expect(t).String("abc123").EqualTo("abc123")
	Expect(t).String("123abc").EqualTo("123abc")
}

func TestStringMatchesPattern(t *testing.T) {
	Expect(t).String("abc123").MatchesPattern("abc[0-9]+")
	Expect(t).String("abc123").MatchesPattern("[a-z]+123")
}
