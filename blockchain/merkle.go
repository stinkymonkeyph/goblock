package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
)

type Node struct {
	Left  *Node
	Right *Node
	Hash  []byte
}

func HashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func HashTransaction(tx *Transaction) []byte {
	txBytes, _ := json.Marshal(tx)
	return txBytes
}

func CreateLeafNodes(transactions []*Transaction) []*Node {
	var leaves []*Node

	for _, tx := range transactions {
		hash := HashTransaction(tx)
		leaves = append(leaves, &Node{Hash: hash})
	}

	return leaves
}

func BuildMerkleTree(leaves []*Node) *Node {
	if len(leaves) == 0 {
		return nil
	}
	for len(leaves) > 1 {
		var newLevel []*Node
		for i := 0; i < len(leaves); i += 2 {
			if i+1 < len(leaves) {
				combinedHash := append(leaves[i].Hash, leaves[i+1].Hash...)
				newNode := &Node{
					Left:  leaves[i],
					Right: leaves[i+1],
					Hash:  HashData(combinedHash),
				}
				newLevel = append(newLevel, newNode)
			} else {
				combinedHash := append(leaves[i].Hash, leaves[i].Hash...)
				newNode := &Node{
					Left:  leaves[i],
					Right: leaves[i],
					Hash:  HashData(combinedHash),
				}
				newLevel = append(newLevel, newNode)
			}
		}
		leaves = newLevel
	}
	return leaves[0]
}

func GenerateMerkleProof(transactions []*Transaction, index int) ([][]byte, error) {
	if index < 0 || index >= len(transactions) {
		return nil, errors.New("invalid index")
	}

	leaves := CreateLeafNodes(transactions)
	proof := make([][]byte, 0)

	for len(leaves) > 1 {
		var newLevel []*Node
		for i := 0; i < len(leaves); i += 2 {
			if i+1 < len(leaves) {
				if i == index || i+1 == index {
					siblingIndex := i ^ 1
					proof = append(proof, leaves[siblingIndex].Hash)
				}
				combinedHash := append(leaves[i].Hash, leaves[i+1].Hash...)
				newNode := &Node{
					Left:  leaves[i],
					Right: leaves[i+1],
					Hash:  HashData(combinedHash),
				}
				newLevel = append(newLevel, newNode)
			} else {
				combinedHash := append(leaves[i].Hash, leaves[i].Hash...)
				newNode := &Node{
					Left:  leaves[i],
					Right: leaves[i],
					Hash:  HashData(combinedHash),
				}
				newLevel = append(newLevel, newNode)
				if i == index {
					proof = append(proof, leaves[i].Hash)
				}
			}
		}
		leaves = newLevel
		index /= 2
	}
	return proof, nil
}

func VerifyTransaction(rootHash []byte, tx *Transaction, proof [][]byte) bool {
	currentHash := HashTransaction(tx)
	for _, siblingHash := range proof {
		combinedHash := append(currentHash, siblingHash...)
		currentHash = HashData(combinedHash)
	}
	return bytes.Equal(currentHash, rootHash)
}
