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

func interpretAddTransactionResult(result blockchain.AddTransactionResult) {
	if result == blockchain.SUCCESS {
		fmt.Println("tx succesfully added to the pool")
	} else if result == blockchain.FAILED_INVALID_SIGNATURE {
		fmt.Println("tx failed to be added to the pool, invalid signature")
	} else if result == blockchain.FAILED_INSUFFICIENT_BALANCE {
		fmt.Println("tx failed to be added to the pool, insufficient balance")
	}
}

func main() {
	minerWallet := wallet.NewWallet()
	bc := blockchain.NewBlockchain(minerWallet.Address(), uint16(3111))

	mnemonic := "canyon wood useful gather anxiety elder stomach kid behind rebel pottery tuition maximum video aisle umbrella slush forward come aware remove guilt olympic hard"

	senderWallet, _ := wallet.ImportWallet(mnemonic)
	receiverWallet := wallet.NewWallet()

	//since transactions will fail because sender wallet don't have balance, we will airdrop coins to sender wallet
	bc.Airdrop(senderWallet.Address())
	bc.Mining()

	t := wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 130.20, blockchain.TRANSFER)

	_, result := bc.AddTransaction(senderWallet.Address(), receiverWallet.Address(), 130.20, senderWallet.PublicKey(), t.GenerateSignature(), blockchain.TRANSFER)

	interpretAddTransactionResult(result)

	bc.Mining()

	t = wallet.NewTransaction(senderWallet.PrivateKey(), senderWallet.PublicKey(), senderWallet.Address(), receiverWallet.Address(), 120, blockchain.TRANSFER)

	_, result = bc.AddTransaction(senderWallet.Address(), receiverWallet.Address(), 120, senderWallet.PublicKey(), t.GenerateSignature(), blockchain.TRANSFER)

	interpretAddTransactionResult(result)

	bc.Mining()

	walletTransactions := bc.GetTransactionsByWalletAddress(senderWallet.Address())

	fmt.Println("Sender Wallet Transctions: ")
	//here we will make test data indexes
	for _, wt := range walletTransactions {
		bh := wt.BlockHeight
		transactionIndex := wt.TransactionIndex
		block, _ := bc.GetBlockByHeight(bh)
		block.Print()
		tx := block.GetTransactionByOrderNumber(transactionIndex)
		tx.Print()
	}

	senderWalletBalance := bc.GetWalletBalanceByAddress(senderWallet.Address())
	fmt.Printf("Sender Balance: \t %1f \n", senderWalletBalance)
}
