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


	// To encode an integer, you can use the EncodeInt which takes an
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
