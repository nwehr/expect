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
    expect.T(t).Int(1).ToEqual(1)
    expect.T(t).Int(5).ToBeGreaterThan(1)
    expect.T(t).Int(1).ToBeLessThan(10)
}

func TestStrings(t *testing.T) {
    // fail with error
    expect.T(t).String("abc123").ToEqual("abc123")
    expect.T(t).String("abc123").ToMatchPattern("abc[0-9]+")
}

func TestNoError(t *testing.T) {
    var err error
    expect.T(t).NoError(err)
}

```
