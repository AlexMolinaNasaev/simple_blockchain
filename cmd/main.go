package main

import (
	"fmt"

	"github.com/alexmolinanasaev/simple_blockchain/pkg/blockchain"
)

func main() {
	fmt.Println("Simple blockchain")

	b := blockchain.GetGenesisBlock()
	fmt.Println(b.CalcHash())
}
