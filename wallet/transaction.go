package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/stinkymonkeyph/goblock/blockchain"
	"github.com/stinkymonkeyph/goblock/utils"
)

type Transaction struct {
	id               [32]byte
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	SenderAddress    string
	RecipientAddress string
	Value            float32
	TransactionType  blockchain.TransactionType
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32, transactionType blockchain.TransactionType) *Transaction {
	id := utils.GenerateTransactionId(sender, recipient, value)
	return &Transaction{id: id, senderPrivateKey: privateKey, senderPublicKey: publicKey, SenderAddress: sender, RecipientAddress: recipient, Value: value, TransactionType: transactionType}
}

func (t *Transaction) GetId() [32]byte {
	return t.id
}

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])

	return &utils.Signature{R: r, S: s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id              string  `json:"id"`
		Sender          string  `json:"sender_address"`
		Recipient       string  `json:"recipient_address"`
		Value           float32 `json:"value"`
		TransactionType string  `json:"transaction_type"`
	}{
		Id:              fmt.Sprintf("%x", t.id),
		Sender:          t.SenderAddress,
		Recipient:       t.RecipientAddress,
		Value:           t.Value,
		TransactionType: t.TransactionType.String(),
	})
}
