package main

func main() {
	bc := NewBlockChain()
	defer bc.db.Close()

	//构造客户端
	cli := CLI{bc}
	cli.Run()
}
