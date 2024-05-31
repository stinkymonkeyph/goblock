package block

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)

	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp\t %d \n", b.timestamp)
	fmt.Printf("nonce\t %d \n", b.nonce)
	fmt.Printf("previousHash\t %s \n", b.previousHash)
	fmt.Printf("transactions\t %s \n", b.transactions)
}
