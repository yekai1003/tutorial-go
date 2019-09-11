package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//区块结构
type Block struct {
	Timestamp     int64  //时间戳
	Data          []byte //模拟交易数据
	PrevBlockHash []byte //前一区块的hash值
	Hash          []byte //本区块hash值
	Nonce         int
}

//构造block的函数
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	//block.SetHash() //设置hash值

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//构造创世块区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//计算哈希值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
