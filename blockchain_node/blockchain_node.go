package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/stinkymonkeyph/goblock/blockchain"
	"github.com/stinkymonkeyph/goblock/wallet"
)

var cache map[string]*blockchain.BlockChain = make(map[string]*blockchain.BlockChain)

type BlockchainNode struct {
	port uint16
}

func NewBlockchainNode(port uint16) *BlockchainNode {
	return &BlockchainNode{port}
}

func (bcn *BlockchainNode) Port() uint16 {
	return bcn.port
}

func (bcn *BlockchainNode) GetChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	bc := bcn.GetBlockchain()
	m, _ := json.Marshal(bc)

	_, err := w.Write(m)

	if err != nil {
		log.Fatal("something went wrong while processing request", err)
	}
}

func (bcn *BlockchainNode) GetBlockchain() *blockchain.BlockChain {
	bc, ok := cache["blockchain"]

	if !ok {
		minerWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minerWallet.Address(), bcn.Port())
		cache["blockchain"] = bc
	}

	return bc
}

func (bcn *BlockchainNode) GetWalletBalanceByAddress(w http.ResponseWriter, r *http.Request) {
	walletAddress := r.PathValue("wallet_address")
	bc := bcn.GetBlockchain()
	balance := bc.GetWalletBalanceByAddress(walletAddress)

	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]float32{"balance": balance})
	if err != nil {
		log.Fatal("something went wrong while processing request", err)
	}
}

func (bcn *BlockchainNode) Run() {
	http.HandleFunc("GET /", bcn.GetChain)
	http.HandleFunc("GET /balance/{wallet_address}", bcn.GetWalletBalanceByAddress)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcn.port)), nil))
}
