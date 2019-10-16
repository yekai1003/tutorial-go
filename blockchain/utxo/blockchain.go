package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

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
func (bc *BlockChain) AddBlock(txs []*Transaction) {
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
	newBlock := NewBlock(txs, lastHash)

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

			coinbasetx := Transaction{}
			genesis := NewGenesisBlock(&coinbasetx)
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

func dbExists() bool {
	_, err := os.Stat(dbFile)
	if err != nil {
		os.IsNotExist(err)
		return false
	}

	return true
}

//创建区块链
func CreateBlockchain(address string) *BlockChain {
	if dbExists() == false {
		fmt.Println("No existing blockchain .create one first")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		tip = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
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

//找到所有未使用的UTXO
func (bc *BlockChain) FindUtxo(address string) []TXOutput {
	// 遍历查找所有的交易
	var UTXOs []TXOutput

	unspentTxs := bc.FindUnspentTransactions(address)
	// 对交易进行一个简单验证 unlocked

	for _, tx := range unspentTxs {
		for _, out := range tx.Vout {
			if out.CanUnlockedOutPutWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

//找到所有的交易
func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction {
	//遍历区块链
	var unspentTXs []Transaction

	spentTXOs := make(map[string][]int)

	bci := bc.Iterator()

	//遍历
	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			//找到和账户相关的
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {

				//去掉无效的
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanUnlockedOutPutWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			//普通交易
			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutPutWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}

		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}

//判断余额是否充足
func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentTXOs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0
Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.CanUnlockedOutPutWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentTXOs[txID] = append(unspentTXOs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentTXOs
}
