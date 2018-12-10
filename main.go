package main

import (
	"MyBlock/models"
)

func init() {

}

func main() {

	//bc := models.NewBlockchain()
	//defer bc.DB.Close()

	cli := models.CLI{}
	cli.Run()

	//bc.AddBlock("Send 1 BTC to Ivan")
	//bc.AddBlock("Send 2 more BTC to Ivan")
	//
	//for _, block := range bc.Blocks {
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	fmt.Printf("Nonce:%d\n", block.Nonce)
	//
	//	pow := models.NewProofOfWork(block)
	//
	//	fmt.Println("Pow:", strconv.FormatBool(pow.ValidateHash()))
	//
	//	fmt.Println()
	//}

}

/*测试一

2018-06-27 17:50:36 Mining the block containing "Gennesis Block"
00000031b90b807c5de5d331026d6ae0bc833aa02f061254e004c3ee1319b6a8
用时: 82.888654558 秒

2018-06-27 17:51:59 Mining the block containing "Send 1 BTC to Ivan"
0000005d69c1eec1e4759bf2ba2551f399211e7ec531d219e87e47e01e14ce4b
用时: 111.248072225 秒

2018-06-27 17:53:50 Mining the block containing "Send 2 more BTC to Ivan"
0000008136e04eff11edde3119459e37745cc1a93cb59de37d6411824690182e
用时: 33.798384443 秒


Prev. hash:
Data: Gennesis Block
Hash: 00000031b90b807c5de5d331026d6ae0bc833aa02f061254e004c3ee1319b6a8
Nonce:19167311
Pow: true

Prev. hash: 00000031b90b807c5de5d331026d6ae0bc833aa02f061254e004c3ee1319b6a8
Data: Send 1 BTC to Ivan
Hash: 0000005d69c1eec1e4759bf2ba2551f399211e7ec531d219e87e47e01e14ce4b
Nonce:24576064
Pow: true

Prev. hash: 0000005d69c1eec1e4759bf2ba2551f399211e7ec531d219e87e47e01e14ce4b
Data: Send 2 more BTC to Ivan
Hash: 0000008136e04eff11edde3119459e37745cc1a93cb59de37d6411824690182e
Nonce:7532677
Pow: true
*/

/*测试二

2018-06-27 17:55:50 Mining the block containing "Gennesis Block"
000000c2cad5a93134ff96ce4ea9f534b06fb781cc0195c8b64511fd112d4bd0
用时: 2.061167316 秒

2018-06-27 17:55:52 Mining the block containing "Send 1 BTC to Ivan"
0000003471415bdeabbf3d5276102b5bf036338549c41efec61c9424c58bd744
用时: 149.88423032 秒

2018-06-27 17:58:22 Mining the block containing "Send 2 more BTC to Ivan"
000000133d3a4f8887845c12d6cfc4c4a03d9acde48b397229426f2a1d19bd2c
用时: 222.602189344 秒


Prev. hash:
Data: Gennesis Block
Hash: 000000c2cad5a93134ff96ce4ea9f534b06fb781cc0195c8b64511fd112d4bd0
Nonce:519102
Pow: true

Prev. hash: 000000c2cad5a93134ff96ce4ea9f534b06fb781cc0195c8b64511fd112d4bd0
Data: Send 1 BTC to Ivan
Hash: 0000003471415bdeabbf3d5276102b5bf036338549c41efec61c9424c58bd744
Nonce:33090994
Pow: true

Prev. hash: 0000003471415bdeabbf3d5276102b5bf036338549c41efec61c9424c58bd744
Data: Send 2 more BTC to Ivan
Hash: 000000133d3a4f8887845c12d6cfc4c4a03d9acde48b397229426f2a1d19bd2c
Nonce:50298517
Pow: true
*/
