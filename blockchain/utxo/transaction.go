package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type TXInput struct {
	Txid      []byte
	Vout      int //要引用的交易输出的编号
	ScriptSig string
}

type TXOutput struct {
	Value        int //比特币数量
	ScriptPubKey string
}

type Transaction struct {
	ID   []byte //交易ID
	Vin  []TXInput
	Vout []TXOutput
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic("Failed to Encode", err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

//判断coinbase交易应该与coinbase交易生成息息相关
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && tx.Vin[0].Vout == -1 && len(tx.Vin[0].Txid) == 0
}

func (in *TXInput) CanUnlockOutPutWith(unlockData string) bool {
	return in.ScriptSig == unlockData
}

func (out *TXOutput) CanUnlockedOutPutWith(unlockData string) bool {
	return out.ScriptPubKey == unlockData
}

//生成coinbase交易
func NewCoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = "coinbase to " + to
	}
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{50, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}

//生成普通交易
func NewUtxoTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	// 验证from 余额是否充足
	// 构造交易 txin txout
	return &Transaction{}
}
