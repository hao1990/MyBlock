package models

import (
	"MyBlock/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"
)

//定义常量 挖矿难度
const targetBits = 24

var (
	maxNonce = math.MaxInt64
)

/*工作证明*/
type ProofOfWork struct {
	block  *Block   //区块指针
	target *big.Int //"目标位"指针
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) //左移256 - targetBits位
	fmt.Println("目标:", target)               //6901746346790563787434755862277025452451108972170386555162524223799296
	pow := &ProofOfWork{b, target}
	return pow
}

/*准备数据
代码只是 把区块的字段和 "目标"、计数器值 合并到一块
*/
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	//nonce 当前计数器
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int //哈希值的整数表示
	var hash [32]byte
	nonce := 0

	startTime := time.Now()

	fmt.Printf("%v Mining the block containing \"%s\"\n", startTime.Format("2006-01-02 15:04:05"), pow.block.Data)

	for nonce < maxNonce {
		//整备数据
		data := pow.prepareData(nonce)

		//用SHA-256计算 哈希值
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)

		//把哈希值转换成大整型
		hashInt.SetBytes(hash[:])

		//与目标进行比较
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n用时:", time.Now().Sub(startTime).Seconds(), "秒")

	return nonce, hash[:]
}

func (pow *ProofOfWork) ValidateHash() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
