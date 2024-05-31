package blockchain

import (
	"fmt"
	"log"
	"strings"

	"github.com/stinkymonkeyph/goblock/block"
)

type BlockChain struct {
	transactionPool []string
	chain           []*block.Block
}

func NewBlockchain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock(0, "hash #0 genesis block")
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *block.Block {
	b := block.NewBlock(nonce, previousHash)
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

func init() {
	log.SetPrefix("Blockchain Node: ")
}
