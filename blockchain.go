package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string, transactions []string) *Block {
	b := new(Block)

	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp\t %d \n", b.timestamp)
	fmt.Printf("nonce\t %d \n", b.nonce)
	fmt.Printf("previousHash\t %s \n", b.previousHash)
	fmt.Printf("transactions\t %s \n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain Node: ")
}

func main() {
	b := NewBlock(0, "root hash", []string{"root tx"})
	b.Print()
}
