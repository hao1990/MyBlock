package models

/*区块链*/
type Blockchain struct {
	Blocks []*Block //区块数组(切片)
}

func (bc *Blockchain) AddBlock(data string) {
	chainLen := len(bc.Blocks)
	//找到上一个区块
	prveBlock := bc.Blocks[chainLen-1]
	//生成新区块
	newBlock := NewBlock(data, prveBlock.Hash)
	//将新区块加入区块链的区块数组中
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	//创世区块
	gennesisBlock := NewGennesisBlock()
	//用创世区块生成区块链
	blockchain := Blockchain{[]*Block{gennesisBlock}}

	return &blockchain
}
