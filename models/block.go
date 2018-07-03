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
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{}, //空的
	}
	//自己生成Hash值
	block.SetHash()

	return block
}

/*给区块添加一个方法，设置区块的哈希值
这里我们仅仅是取区块的字段，并把它们拼接起来，然后技术拼接后数据的SHA256散列值*/
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	//拼接
	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	hash := sha256.Sum256(header)
	//给自己的Hash负值
	b.Hash = hash[:]
}

/*生成 创世区块*/
func NewGennesisBlock() *Block {
	return NewBlock("Gennesis Block", []byte{})
}
