package main

import (
	"fmt"
)

func main() {
	//1. 创建区块链
	bc := NewBlockChain()
	//2. 添加节点
	bc.AddBlock("Send 1 BTC to Yekai")
	bc.AddBlock("Send 2 more BTC to Fuhongxue")
	//3. 区块链浏览
	for _, block := range bc.blocks {
		fmt.Printf("Prev.Hash = %x\n", block.PrevBlockHash)
		fmt.Printf("Data = %s\n", block.Data)
		fmt.Printf("Hash = %x\n", block.Hash)
		fmt.Println()
	}
}
