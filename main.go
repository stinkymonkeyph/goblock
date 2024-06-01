package main

import (
	"log"

	"github.com/stinkymonkeyph/goblock/blockchain"
)

func init() {
	log.SetPrefix("Goblock Node: ")
}

func main() {
	bc := blockchain.NewBlockchain("miner")

	bc.AddTransaction("Rick", "Morty", 130)
	bc.AddTransaction("Gaben", "Notail", 120)

	bc.Mining()

	bc.AddTransaction("Rick", "Morty", 2000)
	bc.AddTransaction("Gaben", "Notail", 1220)

	bc.Mining()

	bc.Print()
}
