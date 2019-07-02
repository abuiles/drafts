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
