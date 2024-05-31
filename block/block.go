package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/stinkymonkeyph/goblock/transaction"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*transaction.Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*transaction.Transaction) *Block {
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
	fmt.Printf("previousHash\t %x \n", b.previousHash)
	fmt.Printf("transaction_count \t %d \n", len(b.transactions))
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64                      `json:"timestamp"`
		Nonce        int                        `json:"nonce"`
		PreviousHash [32]byte                   `json:"previous_hash"`
		Transactions []*transaction.Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}
