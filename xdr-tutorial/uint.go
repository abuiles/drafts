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
