package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stinkymonkeyph/goblock/utils"
)

type Transaction struct {
	SenderAddress    string
	RecipientAddress string
	Value            float32
	Signature        *utils.Signature
}

func NewTransaction(sender string, recipient string, value float32, signature *utils.Signature) *Transaction {
	return &Transaction{SenderAddress: sender, RecipientAddress: recipient, Value: value, Signature: signature}
}

func (t *Transaction) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("sender_address \t%s \n", t.SenderAddress)
	fmt.Printf("recipient_address \t%s \n", t.RecipientAddress)
	fmt.Printf("value \t%1f \n", t.Value)
	if t.SenderAddress == MINING_SENDER {
		fmt.Println("signature miner tx")
	} else {
		fmt.Printf("signature \t%064x%064x \n", t.Signature.R, t.Signature.S)
	}
	fmt.Printf("%s \n", strings.Repeat("-", 40))
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_address"`
		Recipient string  `json:"recipient_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.SenderAddress,
		Recipient: t.RecipientAddress,
		Value:     t.Value,
	})
}
