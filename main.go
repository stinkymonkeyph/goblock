package main

import (
	"log"

	"github.com/stinkymonkeyph/goblock/blockchain"
)

func init() {
	log.SetPrefix("Blockchain Node: ")
}

func main() {
	bc := blockchain.NewBlockchain()
	bc.CreateBlock(23, "hash #1")
	bc.Print()
}
