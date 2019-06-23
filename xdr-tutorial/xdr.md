# XDR for Stellar developers

This article offers an introduction to XDR, which is a standard for
describing and encoding data. Any kind of developer can benefit from
reading this article, but it will be oriented towards Stellar
developers. By the end of the article, you'll be familiar with XDR,
the golang XDR library and understand how XDR is used in the
Stellar ecosystem.

## History

While `JSON` is today the predominant data format for sending messages
around the web, before that, there was `XML` and even before that,
people were using other formats to send messages between machines.

When you are sending messages, you need a way to describe the
information you are sending. Without any context, what are you
supposed to do if you receive a message with the following bytes?

```golang Sequence of bytes represented in decimal
[0, 0, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100]
```

In the mid 1980s Sun Microsystems came up with a proposal which would
allow you not only to send data in `bytes` but also describe the shape
of that data, this format was called External Data Representation
Standard or XDR. [^1]

Following the XDR spec, the `bytes` above represent the string
`hello world`, where the first `4 bytes` tells you the number of
`bytes` in the string (`11` in this case), and the the
rest are the UTF-8 `bytes` representantion of the string. [^2]

XDR is used by a variety of systems like Stellar Payment Network, ZFS
File System and The SpiderMonkey JavaScript engine.

XDR can be compared to Google's [protocol buffers](https://developers.google.com/protocol-buffers/).

In the following sections you'll learn about the relationship between
XDR and Stellar, the basic data types and then explore each data type
while playing with the golang library.

## Stellar and XDR

>XDR, also known as External Data Representation, is used throughout
>the Stellar network and protocol. The ledger, transactions, results,
>history, and even the messages passed between computers running
>stellar-core are encoded using XDR.
>-https://www.stellar.org/developers/guides/concepts/xdr.html

Stellar encodes all its messages using XDR, but you can develop
applications on top of Stellar without worrying about how XDR works,
the SDKs help you interact with the network and convert data back and
forth in XDR format.

The Stellar's guides list some reasons on why XDR was chosen, like:

 - It is very compact, so it can be transmitted quickly and stored with minimal disk space.
 - Data encoded in XDR is reliably and predictably stored. Fields are always in the same order, which makes cryptographically signing and verifying XDR messages simple.
 - XDR definitions include rich descriptions of data types and structures, which is not possible in simpler formats like JSON, TOML, or YAML.

Nicolas Barry also expands on the subject in the following stack  overflow question where someone asked ["Why did the project settle on XDR for external data serialisation?"](https://stellar.stackexchange.com/a/284/1066). It  mentions similar points, but adds emphasis on the extra security added by using the protocol, specifically as the protocol evolves.


## Basic Block Size

XDR requires each data block to have a minimum number of bytes. All
items require a multiple of four bytes (or 32 bits) of data. If the
bytes needed to contain the data are not a multiple of four, then zero
bytes are added to the item to make the total byte count a multiple of
4.

The RFC includes a visual representation of this, which looks like the
following, where each box (delimited by a plus sign at the 4 corners
and vertical bars and dashes) depicts a byte.

        +--------+--------+...+--------+--------+...+--------+
        | byte 0 | byte 1 |...|byte n-1|    0   |...|    0   |   BLOCK
        +--------+--------+...+--------+--------+...+--------+
        |<-----------n bytes---------->|<------r bytes------>|
        |<-----------n+r (where (n+r) mod 4 = 0)>----------->|


Given the restriction above, the following data block which represents a string is not valid, since it contains 15 bytes.

```golang
[0, 0, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100]
```

To make it compliant with XDR you'll need to append an extra `0` after the byte `100`, resulting in `16` bytes.

```golang
[0, 0, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0]
```

Next, let's talk about data types and start playing with them.

## Data types

XDR defines the following data types:

- Integer
- Unsigned Integer
- Enumeration
- Boolean
- Hyper Integer and Unsigned Hyper Integer
- Floating-Point
- Double-Precision Floating-Point
- Quadruple-Precision Floating-Point
- Fixed-Length Opaque Data
- Variable-Length Opaque Data
- String
- Fixed-Length Array
- Variable-Length Array
- Structure
- Discriminated Union
- Void
- Constant
- Typedef
- Optional-Data

We shall start by looking at the most basic data types and then
explore more complex ones like structures and discriminated unions.

## Integer

From the RFC:

   An XDR signed integer is a 32-bit datum that encodes an integer in
   the range [-2147483648,2147483647].  The integer is represented in
   two's complement notation.  The most and least significant bytes are
   0 and 3, respectively.  Integers are declared as follows:

         int identifier;

## Representing data with XDR

```xdr
struct PaymentOp
{
    AccountID destination; // recipient of the payment
    Asset asset;           // what they end up with
    int64 amount;          // amount they end up with
};
```

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


## Footnotes

- [^1] https://tools.ietf.org/html/rfc4506
- [^2] The spec talks about ASCII bytes but current implementations use UTF-8 which has replaced ASCII.
