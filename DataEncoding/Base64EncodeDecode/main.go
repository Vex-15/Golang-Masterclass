package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := "Welcome to Golang Masterclass! Leave a star on my GitHub repo if you like it :D"

	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	println("Encoded Data: ", encodedData)

	decodedData := "V2VsY29tZSB0byBHb2xhbmcgTWFzdGVyY2xhc3MhIExlYXZlIGEgc3RhciBvbiBteSBHaXRIdWIgcmVwbyBpZiB5b3UgbGlrZSBpdCA6RA=="
	decodedStr, err := base64.StdEncoding.DecodeString(decodedData)
	if err != nil {
		log.Fatal(err)
	}
	println("Decoded Data: ", string(decodedStr))

	if string(decodedStr) != data {
		fmt.Println("Decoded data is not same as original data")
	}
}
