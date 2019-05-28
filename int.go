package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/nullstyle/go-xdr/xdr3"
)


func main() {
	version := uint32(10)

	// encode an integer

	var w bytes.Buffer

	enc := xdr.NewEncoder(&w)

	_, err := enc.EncodeUint(version)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("encoded data:", w.Bytes())
	fmt.Println("encoded data:", hex.EncodeToString(w.Bytes()))

	// decode an integer

	dec := xdr.NewDecoder(bytes.NewReader(w.Bytes()))

	version, _, err= dec.DecodeUint()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("decoded data", version)
}
