package main

import (
	"flag"
	"log"

	"github.com/stinkymonkeyph/goblock/blockchain"
	"github.com/stinkymonkeyph/goblock/wallet"
)

func init() {
	log.SetPrefix("Blockchain Node: ")
}

func main() {
	port := flag.Uint("port", 3333, "TCP Port Number for Blockchain Node")
	flag.Parse()

	app := NewBlockchainNode(uint16(*port))
	bc := app.GetBlockchain()
	mnemonic := "canyon wood useful gather anxiety elder stomach kid behind rebel pottery tuition maximum video aisle umbrella slush forward come aware remove guilt olympic hard"

	senderWallet, _ := wallet.ImportWallet(mnemonic)
	receiverWallet := wallet.NewWallet()

	//since transactions will fail because sender wallet don't have balance, we will airdrop coins to sender wallet
	bc.Airdrop(senderWallet.Address())
	bc.Mining()

	t := wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 130.20, blockchain.TRANSFER)

	bc.AddTransaction(t.Id, t.SenderAddress, t.RecipientAddress, t.Value, senderWallet.PublicKey(), t.GenerateSignature(), blockchain.TRANSFER, t.Timestamp)

	bc.Mining()

	t = wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 120, blockchain.TRANSFER)

	bc.AddTransaction(t.Id, t.SenderAddress, t.RecipientAddress, t.Value, senderWallet.PublicKey(), t.GenerateSignature(), blockchain.TRANSFER, t.Timestamp)

	bc.Mining()

	log.Default().Println("Starting Blockchain node on port:", *port)
	app.Run()
}
