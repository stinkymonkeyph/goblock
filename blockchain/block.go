package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	Nonce        int
	PreviousHash [32]byte
	Transactions []*Transaction
	MerkleRoot   []byte
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.Timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.PreviousHash = previousHash
	b.Transactions = transactions
	leafNodes := CreateLeafNodes(transactions)
	fmt.Println(len(leafNodes))
	var merkleRoot []byte
	if len(leafNodes) == 0 {
		merkleRoot = nil
	} else {
		merkleTree := BuildMerkleTree(leafNodes)
		merkleRoot = merkleTree.Hash
	}
	b.MerkleRoot = merkleRoot
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp\t %d \n", b.Timestamp)
	fmt.Printf("nonce\t %d \n", b.Nonce)
	fmt.Printf("previousHash\t %x \n", b.PreviousHash)
	fmt.Printf("transaction_count \t %d \n", len(b.Transactions))
	fmt.Printf("merkle_root \t %x \n", b.MerkleRoot)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) GetTransactionByOrderNumber(orderNumber int) *Transaction {
	return b.Transactions[orderNumber]
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
		MerkleRoot   []byte         `json:"merkle_root"`
	}{
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		PreviousHash: b.PreviousHash,
		Transactions: b.Transactions,
		MerkleRoot:   b.MerkleRoot,
	})
}
