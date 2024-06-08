package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/stinkymonkeyph/goblock/utils"
)

const (
	MINING_DIFICULTY = 4
	MINING_SENDER    = "BLOCKCHAIN REWARD SYSTEM"
	MINING_REWARD    = 1.0
)

type AddTransactionResult int

const (
	SUCCESS AddTransactionResult = iota
	FAILED_INSUFFICIENT_BALANCE
	FAILED_INVALID_SIGNATURE
)

type TransactionBlockHeight struct {
	BlockHeight            int
	TransactionOrderNumber int
	Transaction            *Transaction
}

type WalletTransactionIndex struct {
	transactionBlockHeights []*TransactionBlockHeight
}

type BlockChain struct {
	transactionPool        []*Transaction
	chain                  []*Block
	blockchainAddress      string
	walletTransactionIndex map[string]*WalletTransactionIndex
}

func NewBlockchain(blockchainAddress string) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, b.Hash())
	bc.blockchainAddress = blockchainAddress
	bc.walletTransactionIndex = make(map[string]*WalletTransactionIndex)
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}

	for index, tx := range b.Transactions {
		if _, exists := bc.walletTransactionIndex[tx.SenderAddress]; !exists {
			bc.walletTransactionIndex[tx.SenderAddress] = &WalletTransactionIndex{
				transactionBlockHeights: []*TransactionBlockHeight{},
			}
		}

		if _, exists := bc.walletTransactionIndex[tx.RecipientAddress]; !exists {
			bc.walletTransactionIndex[tx.RecipientAddress] = &WalletTransactionIndex{
				transactionBlockHeights: []*TransactionBlockHeight{},
			}
		}

		blockHeight := len(bc.chain) - 1
		transactionBlockHeight := &TransactionBlockHeight{Transaction: tx, TransactionOrderNumber: index, BlockHeight: blockHeight}
		bc.walletTransactionIndex[tx.SenderAddress].transactionBlockHeights = append(bc.walletTransactionIndex[tx.SenderAddress].transactionBlockHeights, transactionBlockHeight)
		bc.walletTransactionIndex[tx.RecipientAddress].transactionBlockHeights = append(bc.walletTransactionIndex[tx.RecipientAddress].transactionBlockHeights, transactionBlockHeight)

	}

	log.Printf("action=createBlock, status=success, metadata={timestamp: %d, nonce: %d, previousHash: %x} \n", b.Timestamp, b.Nonce, b.PreviousHash)
	return b
}

func (bc *BlockChain) GetBlockByHeight(height int) *Block {
	return bc.chain[height]
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Block %d  %s\n", strings.Repeat("=", 10), i, strings.Repeat("=", 10))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("#", 27))
}

func (bc *BlockChain) LasBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) GetWalletBalance(address string) float32 {
	wtx := bc.GetTransactionsByWalletAddress(address)

	var balance float32

	for _, wt := range wtx {
		if wt.Transaction.SenderAddress == address {
			balance -= wt.Transaction.Value
		}

		if wt.Transaction.RecipientAddress == address {
			balance += wt.Transaction.Value
		}
	}

	return balance
}

func (atr AddTransactionResult) String() string {
	return [...]string{"SUCCESS", "FAILED_INSUFFICIENT_BALANCE", "FAILED_INVALID_SIGNATURE"}[atr]
}

func (bc *BlockChain) AddTransaction(sender string, recipient string, value float32, senderPublicKey *ecdsa.PublicKey, s *utils.Signature, transactionType TransactionType) (bool, AddTransactionResult) {
	t := NewTransaction(sender, recipient, value, s, transactionType)

	if sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
	} else if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		senderWalletBalance := bc.GetWalletBalance(sender)

		if senderWalletBalance < value {
			log.Printf("action=addTransaction, state=failed, status_reason=%s", FAILED_INSUFFICIENT_BALANCE.String())
			return false, FAILED_INSUFFICIENT_BALANCE
		}

		bc.transactionPool = append(bc.transactionPool, t)
	} else {
		log.Printf("action=addTransaction, state=failed, status_reason=%s", FAILED_INVALID_SIGNATURE.String())
		return false, FAILED_INVALID_SIGNATURE
	}

	return true, SUCCESS
}

func (bc *BlockChain) GetTransactionsByWalletAddress(walletAddress string) []*TransactionBlockHeight {
	var transactions []*TransactionBlockHeight
	if wti, exists := bc.walletTransactionIndex[walletAddress]; exists {
		transactions = append(transactions, wti.transactionBlockHeights...)
	}
	return transactions
}

func (bc *BlockChain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *utils.Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *BlockChain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{Timestamp: 0, Nonce: nonce, PreviousHash: previousHash, Transactions: transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *BlockChain) ProofOfWork() int {
	t := bc.CopyTransactionPool()
	previousHash := bc.LasBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, previousHash, t, MINING_DIFICULTY) {
		nonce += 1
	}

	return nonce
}

func (bc *BlockChain) CopyTransactionPool() []*Transaction {
	t := make([]*Transaction, 0)
	for _, tx := range bc.transactionPool {
		t = append(t, NewTransaction(tx.SenderAddress, tx.RecipientAddress, tx.Value, tx.Signature, tx.TransactionType))
	}
	return t
}

func (bc *BlockChain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD, nil, nil, BLOCK_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LasBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *BlockChain) CalculateTotalAmount(address string) float32 {
	var totalAmount float32 = 0
	for _, b := range bc.chain {
		for _, t := range b.Transactions {
			value := t.Value
			if t.RecipientAddress == address {
				totalAmount += value
			}

			if t.SenderAddress == address {
				totalAmount -= value
			}
		}
	}

	return totalAmount
}
