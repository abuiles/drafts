# XDR for Stellar developers

This tutorial introduces XDR which is a standard data serialization
format, for uses such as computer network protocols. It allows data to
be transferred between different kinds of computer systems. Converting
from the local representation to XDR is called encoding. Converting
from XDR to the local representation is called decoding. XDR is
implemented as a software library of functions which is portable
between different operating systems and is also independent of the
transport layer.

## The Spec
## Reimplementing go-xdr
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.

## Integer
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.

```go
  func (enc *Encoder) EncodeInt(v int32) (int, error) {
          var b [4]byte
          b[0] = byte(v >> 24)
          b[1] = byte(v >> 16)
          b[2] = byte(v >> 8)
          b[3] = byte(v)

          n, err := enc.w.Write(b[:])
          if err != nil {
                  msg := fmt.Sprintf(errIOEncode, err.Error(), 4)
                  err := marshalError("EncodeInt", ErrIO, msg, b[:n], err)
                  return n, err
          }

          return n, nil
  }
```

## Unsigned Integer

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.
