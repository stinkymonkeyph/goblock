package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GenerateTransactionId(sender string, recipient string, value float32) [32]byte {
	combined := fmt.Sprintf("%s%s%1f%d", sender, recipient, value, time.Now().UnixNano())
	id := sha256.Sum256([]byte(combined))
	return id
}
