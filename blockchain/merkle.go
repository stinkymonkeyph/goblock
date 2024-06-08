package blockchain

import (
	"crypto/sha256"
	"encoding/json"
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
