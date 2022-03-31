package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error){
	if err != nil{
		log.Panic(err)
	}
}

// encoding strings to bytes
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

// encoding bytes to strings
func FromBytes(i interface{}, data []byte){
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}