package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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

//区块数据序列化
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b) //编码
	if err != nil {
		log.Panic("Failed to Encode ", err)
	}
	return result.Bytes()
}

//解码
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("Failed to Decode ", err)
	}
	return &block
}
