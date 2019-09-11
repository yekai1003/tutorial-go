package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64 //最大范围
)

//挖矿难度
const targetBits = 24

//定义一个PoW结构体
type ProofOfWork struct {
	block  *Block
	target *big.Int //挖矿难度基准值
}

//构造ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) //左移 - 构造挖矿难度基准值

	pow := &ProofOfWork{b, target}

	return pow
}

//根据nonce值重新构造数据

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			Int2Hex(pow.block.Timestamp),
			Int2Hex(int64(targetBits)),
			Int2Hex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

//把函数写入内存

func Int2Hex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num) // 整形数转内存 Big + Endian 大端法
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

//挖矿 PoW
func (pow *ProofOfWork) Run() (int, []byte) {
	fmt.Printf("Begin Mining the block data is %s, maxNoce = %d\n", pow.block.Data, maxNonce)
	nonce := 0
	var hash [32]byte
	var hashInt big.Int
	//循环挖矿
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			//代表符合条件
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n")

	return nonce, hash[:]
}
