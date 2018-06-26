package utils

import (
	"MyBlock/models"
	"math/big"
)

//定义常量 挖矿难度
const targetBits = 24

/*工作证明*/
type ProofOfWork struct {
	block  *models.Block //区块指针
	target *big.Int      //"目标位"指针
}
