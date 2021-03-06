package models

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

/*区块链*/
type Blockchain struct {
	//Blocks []*Block //区块数组(切片)
	tip []byte
	DB  *bolt.DB
}

/*区块迭代器*/
type BlockchainInterator struct {
	currentHash []byte
	db          *bolt.DB
}

//func (bc *Blockchain) AddBlock(data string) {
//	/*
//		chainLen := len(BC.Blocks)
//		//找到上一个区块
//		prveBlock := BC.Blocks[chainLen-1]
//		//生成新区块
//		newBlock := NewBlock(data, prveBlock.Hash)
//		//将新区块加入区块链的区块数组中
//		BC.Blocks = append(BC.Blocks, newBlock)
//	*/
//	var lastHash []byte
//
//	err := bc.DB.View(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		lastHash = b.Get([]byte("l"))
//
//		return nil
//	})
//	if err != nil {
//		log.Panic(err)
//	}
//
//	newBlock := NewBlock(data, lastHash)
//	err = bc.DB.Update(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(blocksBucket))
//		err := b.Put(newBlock.Hash, newBlock.Serialize())
//		if err != nil {
//			log.Panic(err)
//		}
//		err = b.Put([]byte("l"), newBlock.Hash)
//		if err != nil {
//			log.Panic(err)
//		}
//		bc.tip = newBlock.Hash
//		return nil
//	})
//}

func (bc *Blockchain) Iterator() *BlockchainInterator {
	bci := &BlockchainInterator{bc.tip, bc.DB}
	return bci
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
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXOs []Transaction
	spendtTXOs := make(map[string][]int)
	bci := bc.Iterator()
	for {
		block := bci.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		Outputs:
			for outIdx, out := range tx.Vout {
				if spendtTXOs[txID] != nil {

				}
			}
		}
	}
	return unspentTXOs
}

//创建新链的函数
//
func CreateBlockchain(addres string) *Blockchain {
	if dbExists() {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		//创建 第一笔 交易
		cbtx := NewCoinbaseTX(addres, genesisCoinbaseData)
		//创建 创世块
		genesis := NewGennesisBlock(cbtx)

		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesis.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := Blockchain{tip, db}
	return &bc
}
