# XDR for Stellar developers

This article offers an introduction to XDR, which is a standard for
describing and encoding data. The goal of this article is to explore
the standard while recreating the golang library step by step. Any
kind of developer can benefit from reading this article, but it will
be oriented towards Stellar developers. By the end of the article,
you'll be familiar with XDR, the golang XDR library and understand how
XDR is used in the Stellar protocol (ecosystem)?.

XDR uses a language to describe data formats.  The language can only
be used only to describe data; it is not a programming language.

The XDR standard makes the following assumption: that bytes (or
octets) are portable, where a byte is defined to be 8 bits of data.

A given hardware device should encode the bytes onto the various
media in such a way that other hardware devices may decode the bytes
without loss of meaning.  For example, the Ethernet* standard
suggests that bytes be encoded in "little-endian" style [2], or least
significant bit first.

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
