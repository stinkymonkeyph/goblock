package blockchain

import (
	"fmt"
	"log"
	"strings"

	"github.com/stinkymonkeyph/goblock/node/block"
	"github.com/stinkymonkeyph/goblock/node/transaction"
)

const (
	MINING_DIFICULTY = 4
	MINING_SENDER    = "BLOCKCHAIN REWARD SYSTEM"
	MINING_REWARD    = 1.0
)

type BlockChain struct {
	transactionPool   []*transaction.Transaction
	chain             []*block.Block
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *BlockChain {
	b := &block.Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, b.Hash())
	bc.blockchainAddress = blockchainAddress
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *block.Block {
	b := block.NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*transaction.Transaction{}
	log.Printf("action=createBlock, status=success, metadata={timestamp: %d, nonce: %d, previousHash: %x} \n", b.Timestamp, b.Nonce, b.PreviousHash)
	return b
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Block %d  %s\n", strings.Repeat("=", 10), i, strings.Repeat("=", 10))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("#", 27))
}

func (bc *BlockChain) LasBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) AddTransaction(sender string, recipient string, value float32) bool {
	t := transaction.NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
	return true
}

func (bc *BlockChain) ValidProof(nonce int, previousHash [32]byte, transactions []*transaction.Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := block.Block{Timestamp: 0, Nonce: nonce, PreviousHash: previousHash, Transactions: transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	log.Printf("action=validProof, status=verifying, metadata={hash: %s, nonce: %d} \n", guessHashStr, nonce)
	return guessHashStr[:difficulty] == zeros
}

func (bc *BlockChain) ProofOfWork() int {
	t := bc.CopyTransactionPool()
	previousHash := bc.LasBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, previousHash, t, MINING_DIFICULTY) {
		nonce += 1
	}

	return nonce
}

func (bc *BlockChain) CopyTransactionPool() []*transaction.Transaction {
	t := make([]*transaction.Transaction, 0)
	for _, tx := range bc.transactionPool {
		t = append(t, transaction.NewTransaction(tx.SenderAddress, tx.RecipientAddress, tx.Value))
	}
	return t
}

func (bc *BlockChain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LasBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *BlockChain) CalculateTotalAmount(address string) float32 {
	var totalAmount float32 = 0
	for _, b := range bc.chain {
		for _, t := range b.Transactions {
			value := t.Value
			if t.RecipientAddress == address {
				totalAmount += value
			}

			if t.SenderAddress == address {
				totalAmount -= value
			}
		}
	}

	return totalAmount
}
