package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

//帮助函数
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" getbalance -address ADDRESS  -- Get balance of ADDRESS")
	fmt.Println(" send -from FROM -to TO -amount AMOUNT  -- send amount btc to TO from FROM")
	fmt.Println(" printchain  -- print all the blocks of the blockchain")
	fmt.Println(" createblockchain -address ADDRESS  -- create blockchain and send a coinbase")
}

//增加区块
func (cli *CLI) createblockchain(address string) {
	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Sucess!")
}

//遍历区块链
func (cli *CLI) printChain() {
	//得到迭代器
	bci := cli.bc.Iterator()
	//循环next
	for {
		block := bci.Next() //解析当前块+ 设置currenthash=prehash
		fmt.Printf("Prev.Hash = %x\n", block.PrevBlockHash)
		fmt.Printf("Data = %x\n", block.HashTransactions())
		fmt.Printf("Hash = %x\n", block.Hash)
		fmt.Printf("Nonce = %d\n", block.Nonce)
		//验证hash值
		pow := NewProofOfWork(block)
		fmt.Printf("PoW:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevBlockHash) == 0 {
			//已经找到创世块
			break
		}
	}
}

//客户端运行
func (cli *CLI) Run() {
	//判断参数合法
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
	//设置要解析的参数 flag
	createblockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	createblockchain_addr := createblockchainCmd.String("address", "", "coinbase' address")

	//addBlockData := addBlockCmd.String("data", "", "Block data")

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	switch os.Args[1] {
	case "createblockchain":
		err := createblockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic("Failed to createblockchainCmd Parse ", err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic("Failed to printChainCmd Parse ", err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	//如果是增加区块
	if createblockchainCmd.Parsed() {
		//增加区块
		if *createblockchain_addr == "" {
			createblockchainCmd.Usage()
			os.Exit(1)
		}

		cli.createblockchain(*createblockchain_addr)
	}

	//浏览区块
	if printChainCmd.Parsed() {
		cli.printChain()
	}

}
