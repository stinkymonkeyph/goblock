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

	bc.AddTransaction("Rick", "Morty", 130)
	bc.AddTransaction("Gaben", "Notail", 120)

	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce, bc.LasBlock().Hash())
	bc.Print()
}
