package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if err != nil {
		log.Panicf("IntToHex err : %v", err)
	}
	return buffer.Bytes()
}
