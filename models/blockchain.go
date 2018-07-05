package models

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

/*
桶 bucket  > blocks

数据库是价值对存储
key:value
1.
    l:区块链尾端取款的 Hash
2.
	hash1 : 区块1
	hash2 : 区块2

*/
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
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = bucket.Put([]byte("l"), newBlock.Hash)
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

/*
1.打开一个数据库文件。
2.检查里面是否存了区块链。
3.如果里面有：
	1.创建一个新的Blockchain实例。
	2.把Blockchain实例的末端设为数据库中存储的最后一个区块
4.如果没有区块链：
	1.创建创世区块
	2.存入数据库
	3.把创世区块的哈希存作为最后区块哈希存储
	4.创建一个尾部指向创世区块的Blockchain实例
*/
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
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
			tip = bucket.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{tip, db}

	return &blockchain
}
