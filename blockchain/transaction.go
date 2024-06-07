package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stinkymonkeyph/goblock/utils"
)

type TransactionType int

const (
	TRANSFER TransactionType = iota
	BLOCK_REWARD
)

type Transaction struct {
	SenderAddress    string
	RecipientAddress string
	Value            float32
	Signature        *utils.Signature
	TransactionType  TransactionType
}

func (tt TransactionType) String() string {
	return [...]string{"TRANSFER", "BLOCK_REWARD"}[tt]
}

func NewTransaction(sender string, recipient string, value float32, signature *utils.Signature, transactionType TransactionType) *Transaction {
	return &Transaction{SenderAddress: sender, RecipientAddress: recipient, Value: value, Signature: signature, TransactionType: transactionType}
}

func (t *Transaction) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("sender_address \t%s \n", t.SenderAddress)
	fmt.Printf("recipient_address \t%s \n", t.RecipientAddress)
	fmt.Printf("value \t%1f \n", t.Value)
	fmt.Printf("Type \t%s \n", t.TransactionType.String())
	if t.Signature != nil {
		fmt.Printf("signature \t%064x%064x \n", t.Signature.R, t.Signature.S)
	}

	fmt.Printf("%s \n", strings.Repeat("-", 40))
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender          string          `json:"sender_address"`
		Recipient       string          `json:"recipient_address"`
		Value           float32         `json:"value"`
		TransactionType TransactionType `json:"transaction_type"`
	}{
		Sender:          t.SenderAddress,
		Recipient:       t.RecipientAddress,
		Value:           t.Value,
		TransactionType: t.TransactionType,
	})
}
