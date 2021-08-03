## Expect

An assertion library for Go unit testing.

## Example

```
package my_package

import "testing"

func TestInts(t *testing.T) {
    Expect(t).Int(1).ToEqual(1)
    Expect(t).Int(5).ToBeGreaterThan(1)
    Expect(t).Int(1).ToBeLessThan(10)
}

func TestStrings(t *testing.T) {
    Expect(t).String("abc123").ToEqual("abc123")
    Expect(t).String("abc123").ToMatchPattern("abc[0-9]+")
}

```
