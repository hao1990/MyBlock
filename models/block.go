package models

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/*区块*/
type Block struct {
	Timestamp     int64  //区块的创建时间
	Data          []byte //区块中有价值的数据
	PrevBlockHash []byte //上一个区块的散列值
	Hash          []byte //该区块的散列值
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}

/*生成 创世区块*/
func NewGennesisBlock() *Block {
	return NewBlock("Gennesis Block", []byte{})
}
