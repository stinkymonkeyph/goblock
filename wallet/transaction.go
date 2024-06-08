package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"

	"github.com/stinkymonkeyph/goblock/blockchain"
	"github.com/stinkymonkeyph/goblock/utils"
)

type Transaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	SenderAddress    string
	RecipientAddress string
	Value            float32
	TransactionType  blockchain.TransactionType
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32, transactionType blockchain.TransactionType) *Transaction {
	return &Transaction{senderPrivateKey: privateKey, senderPublicKey: publicKey, SenderAddress: sender, RecipientAddress: recipient, Value: value, TransactionType: transactionType}
}

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])

	return &utils.Signature{R: r, S: s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender          string  `json:"sender_address"`
		Recipient       string  `json:"recipient_address"`
		Value           float32 `json:"value"`
		TransactionType string  `json:"transaction_type"`
	}{
		Sender:          t.SenderAddress,
		Recipient:       t.RecipientAddress,
		Value:           t.Value,
		TransactionType: t.TransactionType.String(),
	})
}
