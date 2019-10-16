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
	fmt.Println(" addblock -data BLOCK_DATA  -- add a block to the blockchain")
	fmt.Println(" printchain  -- print all the blocks of the blockchain")
}

//增加区块
func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
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
		fmt.Printf("Data = %s\n", block.Data)
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
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic("Failed to addBlockCmd Parse ", err)
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
	if addBlockCmd.Parsed() {
		//增加区块
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}

		cli.addBlock(*addBlockData)
	}

	//浏览区块
	if printChainCmd.Parsed() {
		cli.printChain()
	}

}
