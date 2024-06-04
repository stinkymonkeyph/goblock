package main

import (
	"log"

	"github.com/stinkymonkeyph/goblock/blockchain"
	"github.com/stinkymonkeyph/goblock/wallet"
)

func init() {
	log.SetPrefix("Goblock Node: ")
}

func main() {
	minerWallet := wallet.NewWallet()
	bc := blockchain.NewBlockchain(minerWallet.Address())

	senderWallet := wallet.NewWallet()
	receiverWallet := wallet.NewWallet()

	t := wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 130)

	bc.AddTransaction(senderWallet.Address(), receiverWallet.Address(), 130, senderWallet.PublicKey(), t.GenerateSignature())

	bc.Mining()

	bc.Print()
}
