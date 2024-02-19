package block

import (
	"blockChain/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func MineBlock(block *models.DbBlock, difficulty int) (*models.DbBlock, error) {
	nonce := 0
	for {
		// Calculate hash
		hash := calculateHash(block)
		leadingZeros := countLeadingZeros(hash)
		// Check if hash meets the required difficulty
		if leadingZeros >= difficulty {
			block.Hash = hash
			return block, nil
		}

		nonce++
		block.Nonce = nonce
	}
}

func calculateHash(block *models.DbBlock) string {
	blockBytes := []byte(block.Id.String() + block.PrevHash + string(block.Data) + strconv.Itoa(block.Nonce))
	hashBytes := sha256.Sum256(blockBytes)
	return hex.EncodeToString(hashBytes[:])
}

func countLeadingZeros(hash string) int {
	count := 0
	for _, char := range hash {
		if char == '0' {
			count++
		} else {
			break
		}
	}
	return count
}

func SerilizeData(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func DeserilizeData(s string) string {
	var response string
	_, err := fmt.Sscanf(s, "%x", &response)
	if err != nil {
		fmt.Println("Error reverting:", err)
	}
	return response
}
