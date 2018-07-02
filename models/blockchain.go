package models

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

/*区块链*/
type Blockchain struct {
	//Blocks []*Block //区块数组(切片)
	tip []byte
	DB  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	/*
		chainLen := len(BC.Blocks)
		//找到上一个区块
		prveBlock := BC.Blocks[chainLen-1]
		//生成新区块
		newBlock := NewBlock(data, prveBlock.Hash)
		//将新区块加入区块链的区块数组中
		BC.Blocks = append(BC.Blocks, newBlock)
	*/
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}

func (bc *Blockchain) Iterator() *BlockchainInterator {
	bci := &BlockchainInterator{bc.tip, bc.DB}
	return bci
}

/*区块迭代器*/
type BlockchainInterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bci *BlockchainInterator) Next() *Block {
	var block *Block
	err := bci.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encoderBlcok := b.Get(bci.currentHash)
		block = DeserializeBlock(encoderBlcok)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bci.currentHash = block.PrevBlockHash
	return block
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("没有发现区块链 》》》 创建新的")
			//创世区块
			gennesisBlock := NewGennesisBlock()
			//用创世区块生成区块链
			//blockchain := Blockchain{[]*Block{gennesisBlock}}

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(gennesisBlock.Hash, gennesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("l"), gennesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = gennesisBlock.Hash

		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{tip, db}

	return &blockchain
}
