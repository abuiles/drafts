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
