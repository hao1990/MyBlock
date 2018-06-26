package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	testStr := "I like donutes"
	nonce := 0
	for nonce < 10 {
		nonce++
		temp := fmt.Sprintf(testStr, nonce)
		byteslice := []byte(temp)
		fmt.Println(sha256.Sum256(byteslice))
	}
}
