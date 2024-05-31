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
	bc.CreateBlock(1, bc.LasBlock().Hash())
	bc.Print()
}
