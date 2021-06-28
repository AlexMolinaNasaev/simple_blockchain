package main

import (
	"fmt"
	"log"

	"github.com/alexmolinanasaev/simple_blockchain/pkg/blockchain"
)

func main() {
	fmt.Println("Simple blockchain")

	chain := blockchain.NewChain(255)
	err := chain.Validate()
	if err != nil {
		log.Fatal(err.Error())
	}

	chain.AddBlock("hello!")
	chain.AddBlock("world!")
	chain.AddBlock("test!")
	chain.AddBlock("this is a transaction")
	chain.AddBlock("here!")
	chain.AddBlock("foobarbaz")

	chain.Blocks[2].Hash = "a81c83a869fc51a4771c18601c0070d5dc714004e9c3c5f6eb5fbc6a2990f4f5"

	err = chain.Validate()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, b := range chain.Blocks {
		fmt.Printf("Block[%d] hash: %s\n", b.Number, b.Hash)
	}
}
