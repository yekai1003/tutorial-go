package main

type BlockChain struct {
	blocks []*Block
}

//新建一个区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//添加一个区块
func (bc *BlockChain) AddBlock(data string) {
	//拿到前一区块的hash值
	prevBlock := bc.blocks[len(bc.blocks)-1]
	//创建新区块
	newBlock := NewBlock(data, prevBlock.Hash)
	//追加到区块链上
	bc.blocks = append(bc.blocks, newBlock)
}
