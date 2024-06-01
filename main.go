package main

import (
	"fmt"
	"log"

	blockchain "github.com/stinkymonkeyph/goblock/node/blockhain"
)

func init() {
	log.SetPrefix("Goblock Node: ")
}

func main() {
	minerAddress := "miner"
	bc := blockchain.NewBlockchain(minerAddress)

	bc.AddTransaction("Rick", "Morty", 130)
	bc.AddTransaction("Gaben", "Notail", 120)

	bc.Mining()

	bc.AddTransaction("Rick", "Morty", 2000)
	bc.AddTransaction("Gaben", "Notail", 1220)

	bc.Mining()

	bc.Print()

	fmt.Printf("Miner Total Value: %1.f \n", bc.CalculateTotalAmount(minerAddress))
}
