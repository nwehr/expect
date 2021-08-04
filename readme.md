## Expect

An assertion library for Go unit testing.

## Example

```
package my_package

import (
    "testing"
    "github.com/nwehr/expect"
)

func TestInts(t *testing.T) {
    // fatally fail
    expect.Fatal(t).Int(1).ToEqual(1)
    expect.Fatal(t).Int(5).ToBeGreaterThan(1)
    expect.Fatal(t).Int(1).ToBeLessThan(10)
}

func TestStrings(t *testing.T) {
    // fail with error
    expect.Error(t).String("abc123").ToEqual("abc123")
    expect.Error(t).String("abc123").ToMatchPattern("abc[0-9]+")
}

```
