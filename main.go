package main

import (
	"MyBlock/models"
	"fmt"
)

func init() {

}

func main() {

	bc := models.NewBlockchain() //会自动生成创世区块

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

/*

Prev.hash:
Data: Gennesis Block
Hash: d8c83ede239641fa29465d8e5e6e0b22dbacd339d60c37978e62379fcd757fa8

Prev.hash: d8c83ede239641fa29465d8e5e6e0b22dbacd339d60c37978e62379fcd757fa8
Data: Send 1 BTC to Ivan
Hash: d1c0fea3cd95b1151fcc0fe530d64ffcdd67ee132deca5eb47704bedded0348b

Prev.hash: d1c0fea3cd95b1151fcc0fe530d64ffcdd67ee132deca5eb47704bedded0348b
Data: Send 2 more BTC to Ivan
Hash: 01edf9d9e00bb3c787ee95b6c125d515e51b5cd2159080e60bbc36c65fb083e6

*/
