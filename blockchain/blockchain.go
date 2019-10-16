package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

//go get -u github.com/boltdb/bolt
// type BlockChain struct {
// 	blocks []*Block
// }

const dbFile = "yekai.db"
const blockBucket = "blocks"

type BlockChain struct {
	tip []byte   //记录当前区块的hash值
	db  *bolt.DB //bolt数据库
}

//定义迭代器结构
type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

//添加一个区块
func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte //当前区块hash值
	//判断是否有区块

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l")) //put - get ===> key === lastHash => l
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	//执行到这里，一定会创建新区块
	newBlock := NewBlock(data, lastHash)

	//修改数据库文件
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize()) // hash=>data
		if err != nil {
			log.Panic("Failed to Put", err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic("Failed to Put:l", err)
		}
		bc.tip = newBlock.Hash

		return nil
	})
}

//启动区块链-创世块

func NewBlockChain() *BlockChain {
	var tip []byte
	//打开数据库文件
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic("Failed to Open ", dbFile, err)
	}
	// update
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket)) //得到bucket
		if b == nil {
			//初始化
			fmt.Println("No existing blockchain found.Create now ...")
			genesis := NewGenesisBlock()
			//创建bucket
			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("Failed to CreateBucket", err)
			}

			// hash=>data  l ==> hash
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic("Failed to Put genesis", err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic("Failed to Put genesis.hash", err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Panic("Failed to update genesis", err)
	}

	bc := BlockChain{tip, db}

	return &bc
}

//得到迭代器结构对象

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.tip, bc.db}
}

//得到下一个结构
func (i *BlockChainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.currentHash)
		block = Deserialize(encodedBlock)
		return nil
	})

	if err != nil {
		log.Panic("Failed to View -->Next", err)
	}

	i.currentHash = block.PrevBlockHash // 指向之前块的hash

	return block
}
