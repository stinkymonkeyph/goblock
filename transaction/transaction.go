package transaction

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("sender_address \t%s \n", t.senderAddress)
	fmt.Printf("recipient_address \t%s \n", t.recipientAddress)
	fmt.Printf("value \t%1f \n", t.value)
	fmt.Printf("%s \n", strings.Repeat("-", 40))
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_address"`
		Recipient string  `json:"recipient_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderAddress,
		Recipient: t.recipientAddress,
		Value:     t.value,
	})
}
