package models

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
}

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchain(address)
	bc.DB.Close()

	fmt.Println("Done!")
}

func (cli *CLI) getBalance(address string) {
	//bc := NewBlockchain(address)
	//defer bc.DB.Close()
	//
	//balance := 0
	//UTXOs := bc.FindUTXO(address)
	//
	//for _, out := range UTXOs {
	//	balance += out.Value
	//}
	//
	//fmt.Printf("Balance of '%s': %d\n", address, balance)
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		os.Exit(100)
	}
}

//func (cli *CLI) addBlock(data string) {
//	cli.BC.AddBlock(data)
//	fmt.Println("Success!")
//}
func (cli *CLI) printChain() {
	//bci := cli.BC.Iterator()
	//
	//for {
	//	block := bci.Next()
	//
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//
	//	pow := NewProofOfWork(block)
	//	fmt.Printf("Pow:%s\n", strconv.FormatBool(pow.ValidateHash()))
	//	fmt.Println()
	//
	//	if len(block.PrevBlockHash) == 0 {
	//		break
	//	}
	//}
}

func (cli *CLI) send(from, to string, amount int) {
	//bc := NewBlockchain(from)
	//defer bc.DB.Close()
	//
	//tx := NewUTXOTransaction(from, to, amount, bc)
	//bc.MineBlock([]*Transaction{tx})
	//fmt.Println("Success!")
}

func (cli *CLI) Run() {
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)

	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

}
