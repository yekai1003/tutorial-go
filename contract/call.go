package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetFileName(address, dirname string) (string, error) {

	data, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("read dir err", err)
		return "", err
	}
	for _, v := range data {
		if strings.Index(v.Name(), address) > 0 {
			//代表找到文件
			return v.Name(), nil
		}
	}

	return "", nil
}

//设置签名
func MakeAuth(addr, pass string) (*bind.TransactOpts, error) {
	keystorePath := "/Users/yekai/eth/data/keystore"
	fileName, err := GetFileName(string([]rune(addr)[2:]), keystorePath)
	if err != nil {
		fmt.Println("failed to GetFileName", err)
		return nil, err
	}

	file, err := os.Open(keystorePath + "/" + fileName)
	if err != nil {
		fmt.Println("failed to open file ", err)
		return nil, err
	}
	auth, err := bind.NewTransactor(file, pass)
	if err != nil {
		fmt.Println("failed to NewTransactor  ", err)
		return nil, err
	}
	return auth, err
}
func main() {
	fmt.Println("hello world")
	cli, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Println("Failed to Dial ", err)
		return
	}
	ins, err := NewPerson(common.HexToAddress("0x27cc9f5e4585cd3ba03df85f96725ef8ba128db1"), cli)
	if err != nil {
		fmt.Println("Failed to NewPerson", err)
	}
	// name, err := ins.Name(nil)
	// fmt.Println("name ==", name)
	auth, err := MakeAuth("0x2ae72cb02aef1322ef6dfcd18577747969680c10", "123")
	if err != nil {
		fmt.Println("failed to MakeAuth")
		return
	}
	tx, err := ins.SetName(auth, "yekai")
	if err != nil {
		fmt.Println("Failed to setname", err)
		return
	}
	fmt.Println(tx.Hash().Hex())
}
