package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

/*区块*/
type Block struct {
	Timestamp    int64          //区块的创建时间
	Transactions []*Transaction //交易
	//Data          []byte //区块中有价值的数据
	PrevBlockHash []byte //上一个区块的散列值
	Hash          []byte //该区块的散列值
	Nonce         int    //计数器值
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	//block.SetHash()

	return block
}

//func (b *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
//	hash := sha256.Sum256(header)
//	b.Hash = hash[:]
//}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

/*生成 创世区块*/
func NewGennesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

/*序列化*/
/*把Blockc 序列化编码成字节数组*/
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

/*把字节数组反序列化成 Block*/
func DeserializeBlock(by []byte) *Block {
	var block Block
	docoder := gob.NewDecoder(bytes.NewReader(by))
	err := docoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
