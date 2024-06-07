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

	t := wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 130, blockchain.TRANSFER)

	bc.AddTransaction(senderWallet.Address(), receiverWallet.Address(), 130, senderWallet.PublicKey(), t.GenerateSignature(), blockchain.TRANSFER)

	bc.Mining()

	t = wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 120, blockchain.TRANSFER)

	bc.AddTransaction(senderWallet.Address(), receiverWallet.Address(), 120, senderWallet.PublicKey(), t.GenerateSignature(), blockchain.TRANSFER)

	bc.Mining()

	walletTransactions := bc.GetTransactionsByWalletAddress(senderWallet.Address())

	fmt.Println("Sender Wallet Transctions: ")
	for _, wt := range walletTransactions {
		bh := wt.BlockHeight
		transactionOrderNumber := wt.TransactionOrderNumber
		block := bc.GetBlockByHeight(bh)
		block.Print()
		tx := block.GetTransactionByOrderNumber(transactionOrderNumber)
		tx.Print()
	}

}
