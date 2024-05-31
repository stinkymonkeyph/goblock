package blockchain

import (
	"fmt"
	"strings"

	"github.com/stinkymonkeyph/goblock/block"
	"github.com/stinkymonkeyph/goblock/transaction"
)

type BlockChain struct {
	transactionPool []*transaction.Transaction
	chain           []*block.Block
}

func NewBlockchain() *BlockChain {
	b := &block.Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *block.Block {
	b := block.NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)

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
