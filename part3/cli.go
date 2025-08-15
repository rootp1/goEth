package main
import "fmt"

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("block added")
}

