package models

import "fmt"

/*在比特币中，实际并没有存储这个数字，
而是基于区块总数进行计算而得：区块总数除以 210000 就是 subsidy。
挖出创世块的奖励是 50 BTC，每挖出 210000 个块后，奖励减半。
在我们的实现中，这个奖励值将会是一个常量（至少目前是）。*/
const subsidy = 10 //奖励的数额

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXInput struct {
	Txid      []byte //这笔交易的id
	Vout      int    //输出的索引
	ScriptSig string //脚本
}

type TXOutput struct {
	Value        int //value 存储的 satoshi的数量,一个satoshi = 百万分之一各 比特币
	ScriptPubKey string
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	return &tx
}
