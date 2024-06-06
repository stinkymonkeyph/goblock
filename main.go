package main

import (
	"fmt"
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

	mnemonic := "canyon wood useful gather anxiety elder stomach kid behind rebel pottery tuition maximum video aisle umbrella slush forward come aware remove guilt olympic hard"

	senderWallet, _ := wallet.ImportWallet(mnemonic)
	receiverWallet := wallet.NewWallet()

	fmt.Printf("sender wallet seed: %s \n", senderWallet.Mnemonic())

	t := wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 130)

	bc.AddTransaction(senderWallet.Address(), receiverWallet.Address(), 130, senderWallet.PublicKey(), t.GenerateSignature())

	bc.Mining()

	t = wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 120)

	bc.AddTransaction(senderWallet.Address(), receiverWallet.Address(), 120, senderWallet.PublicKey(), t.GenerateSignature())

	bc.Mining()

	bc.Print()
}
