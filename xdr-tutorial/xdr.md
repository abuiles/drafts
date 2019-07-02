# XDR for Stellar developers

This article offers an introduction to XDR, which is a standard for
describing and encoding data. Any kind of developer can benefit from
reading this article, but it will be oriented towards Stellar
developers. By the end of the article, you'll be familiar with XDR,
the `golang` XDR library and understand how XDR is used in the
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
while playing with the `golang` library.

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

>An XDR signed integer is a 32-bit datum that encodes an integer in
>the range [-2147483648,2147483647].  The integer is represented in
>two's complement notation.  The most and least significant bytes are
>0 and 3, respectively.
>
>Integers are declared as follows:
>
>         int identifier;
>
> https://tools.ietf.org/html/rfc4506#section-4.1

As described above, signed integers in XDR follow a 32-bit
architecture representation. Assuming you have `golang` installed,
let's play with the `XDR` go library to encode and decode some
integers, and print their representation hex and decimal.

Create a file called `int.go` and then copy the following content:

```golang
// You can run directly from the golang playground https://play.golang.org/p/ArVhlyO8udz

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/nullstyle/go-xdr/xdr3"
)

func encodeAndDecode(anInteger int32) {
	fmt.Println("Encoding and decoding:", anInteger)

	// Use Buffer to write bytes
	var buffer bytes.Buffer

	// The XDR library has a built-in encoder which helps you convert
	// golang types to XDR types. You can create a new encoder with
	// the function `NewEncoder` which takes a buffer
	enc := xdr.NewEncoder(&buffer)


	// To encode an integer, you can use the `EncodeInt` which takes an
	// int32 integer and writes its bytes representation to the
	// buffer
	_, err := enc.EncodeInt(anInteger)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("encoded data in decimal:", buffer.Bytes())
	fmt.Println("encoded data in hex:", hex.EncodeToString(buffer.Bytes()))

	// decode an integer

	// Similar to encoding, you can create a new decoder by using
	// `NewDecoder` and pass a byte stream, in this example, you are
	// passing the buffer with an integer in bytes
	dec := xdr.NewDecoder(bytes.NewReader(buffer.Bytes()))

	// `DecodeInt` reads the content in the buffer passed to
	// `NewDecoder`, it returns its int32 representation
	decoded, _, err := dec.DecodeInt()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("decoded data", decoded)
}

func main() {
	encodeAndDecode(int32(10))
	encodeAndDecode(int32(-10))
}
```

Run the program with `go run int.go` and you should get the following output:

```golang
Encoding and decoding: 10
encoded data in decimal: [0 0 0 10]
encoded data in hex: 0000000a
decoded data 10
Encoding and decoding: -10
encoded data in decimal: [255 255 255 246]
encoded data in hex: fffffff6
decoded data -10
```

In the example above, you created a buffer, a decoder/encoder and then
called the method to handle the data type which you were marshalling,
`enc.EncodeInt` and `enc.DecodeInt`. There is an easier way to do the
same thing using the helper methods `xdr.Marshal` and
`xdr.Unmarshal`. The code above could be rewritten as you are about to
see. To keep the code short, it doesn't include any error handling.

```golang
// You can run directly from the golang playground https://play.golang.org/p/OoVTu_li7LO

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/nullstyle/go-xdr/xdr3"
)

func encodeAndDecode(anInteger int32) {
	fmt.Println("Encoding and decoding:", anInteger)
	var buffer bytes.Buffer

	// Pass a buffer and a value, `Marshal` will use reflection to
	// identify the right encoder.
	xdr.Marshal(&buffer, anInteger)

	fmt.Println("encoded data in decimal:", buffer.Bytes())
	fmt.Println("encoded data in hex:", hex.EncodeToString(buffer.Bytes()))

	var decoded int32

	// Like `Marshal`, `Unmarshal` will use reflection to
	// identify the right decoder.
	xdr.Unmarshal(bytes.NewReader(buffer.Bytes()), &decoded)

	fmt.Println("decoded data", decoded)
}

func main() {
	encodeAndDecode(int32(10))
	encodeAndDecode(int32(-10))
}
```

Running the code above, will yield the same result as the original implementation:

```golang
Encoding and decoding: 10
encoded data in decimal: [0 0 0 10]
encoded data in hex: 0000000a
decoded data 10
Encoding and decoding: -10
encoded data in decimal: [255 255 255 246]
encoded data in hex: fffffff6
decoded data -10
```

From now on, we'll use `Marshal` and `Unmarshal` to convert between
`golang` and `XDR` data types.

## Unsigned Integer


>An XDR unsigned integer is a 32-bit datum that encodes a non-negative
>integer in the range [0,4294967295].  It is represented by an
>unsigned binary number whose most and least significant bytes are 0
>and 3, respectively.
>
>An unsigned integer is declared as follows:
>
>
>         unsigned int identifier;
>
>https://tools.ietf.org/html/rfc4506#section-4.2

Unsigned integers are very similar to integers. The following code shows you how to decode and encode unsigned integers:

```golang
// You can run directly from the golang playground https://play.golang.org/p/M4H7YfgB4yW

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/nullstyle/go-xdr/xdr3"
)

func encodeAndDecode(anInteger uint32) {
	fmt.Println("Encoding and decoding:", anInteger)
	var buffer bytes.Buffer
	xdr.Marshal(&buffer, anInteger)

	fmt.Println("encoded data in decimal:", buffer.Bytes())
	fmt.Println("encoded data in hex:", hex.EncodeToString(buffer.Bytes()))

	var decoded uint32
	xdr.Unmarshal(bytes.NewReader(buffer.Bytes()), &decoded)

	fmt.Println("decoded data", decoded)
}

func main() {
	encodeAndDecode(uint32(1000))
	encodeAndDecode(uint32(0))
}
```

## Enumeration

>Enumerations have the same representation as signed integers.
>Enumerations are handy for describing subsets of the integers.
>Enumerated data is declared as follows:
>
>      enum { name-identifier = constant, ... } identifier;
>
>https://tools.ietf.org/html/rfc4506#section-4.3

Enum allow us to define identifiers for signed integers, internally
they are also stored using 4 bytes. Using the XDR language you can
define an enum like the following:

```xdr
enum OperationType
{
    CREATE_ACCOUNT = 0,
    PAYMENT = 1,
};
```

The following implements the `enum` above in `golang`, notice how we use
the interface `Enum` define in the `XDR` library to force validation
at encode and decode time.

```go
// You can run directly from the golang playground https://play.golang.org/p/X7WuzuDZ8zM

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/nullstyle/go-xdr/xdr3"
)

// Define enum
type OperationType int32

// Possible values for OperationType
const (
	OperationTypeCreateAccount          OperationType = 0
	OperationTypePayment                OperationType = 1
)

// This is used latter to define if the value is valid or not.
var operationTypeMap = map[int32]string{
	0:  "OperationTypeCreateAccount",
	1:  "OperationTypePayment",
}

// ValidEnum validates a proposed value for this enum. Implements the
// Enum interface for OperationType.

// The enum interface is defined here
//
// https://github.com/nullstyle/go-xdr/blob/master/xdr3/main.go#L7
//
// Its used at encode/decode time see the following link:
//
// https://github.com/nullstyle/go-xdr/blob/master/xdr3/encode.go#L658
//
func (e OperationType) ValidEnum(v int32) bool {
	_, ok := operationTypeMap[v]
	return ok
}

func encodeAndDecode(enumValue OperationType) {
	fmt.Println("Encoding and decoding:", enumValue)

	var buffer bytes.Buffer

	_, err := xdr.Marshal(&buffer, enumValue)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("encoded data in decimal:", buffer.Bytes())
	fmt.Println("encoded data in hex:", hex.EncodeToString(buffer.Bytes()))

	var decoded OperationType
	xdr.Unmarshal(bytes.NewReader(buffer.Bytes()), &decoded)

	fmt.Println("decoded data", decoded)
}

func main() {
	encodeAndDecode(OperationType(0))
	encodeAndDecode(OperationType(1))
	encodeAndDecode(OperationType(2))
	encodeAndDecode(OperationType(3))
}
```

Running the code above will produce:

```golang
Encoding and decoding: 0
encoded data in decimal: [0 0 0 0]
encoded data in hex: 00000000
decoded data 0
Encoding and decoding: 1
encoded data in decimal: [0 0 0 1]
encoded data in hex: 00000001
decoded data 1
Encoding and decoding: 2
xdr:encode: invalid enum - value: '2'
Encoding and decoding: 3
xdr:encode: invalid enum - value: '3'
```


## Boolean, Hyper Integer and Unsigned Hyper Integer, Floating-Point, Double-Precision Floating-Point

By now you must see a pattern on how to encode/decode some data types
between XDR and `golang`, we won't include examples for the ones listed
above since it is pretty much the same code as the previous
examples. We'll list the mapping type in go and the number of bytes
used to stored each type:

- Boolean:  `bool` - 32-bits (4-bytes).
- Hyper Integer and Unsigned Hyper Integer: `int64` and `uint64` - 64-bits (8-bytes).
- Floating-Point: `float32` - 32-bits (4-bytes).
- Double-Precision Floating-Point: `float64` - 64-bits (8-bytes).

##  Quadruple-Precision Floating-Point

>The standard defines the encoding for the quadruple-precision
>floating-point data type "quadruple" (128 bits or 16 bytes).
>
>https://tools.ietf.org/html/rfc4506#section-4.8

This is supported by the standard but not implemented in the golang library.

## Fixed-Length Opaque Data
## Variable-Length Opaque Data
## String
## Fixed-Length Array
## Variable-Length Array
## Structure
## Discriminated Union
## Void
## Constant
## Typedef
## Optional-Data

## Footnotes

- [^1] https://tools.ietf.org/html/rfc4506
- [^2] The spec talks about ASCII bytes but current implementations use UTF-8 which has replaced ASCII.
