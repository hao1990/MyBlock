package models

/*区块链
区块链本质上是有特定结构的数据库：有序的、向后链表。
这意味着区块按照插入顺序排列并且每个区块连接着上一区块。
通过这种结构可以迅速地获取到最近一个区块，也可以通过哈希值（高效地）获取一个区块。
*/

/*这一节暂不涉及数据库,暂用一个数组实现,数组里按顺序存"区块"*/
/*声明一个区块链结构或者类*/
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
