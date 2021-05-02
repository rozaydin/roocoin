package bip39

import (
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
	"time"
)

func NewMnemonic() []string {

	entrophy := generateEntrophy(32)
	entrophyLen := len(entrophy)
	checksum := calculateSha256(entrophy)
	checksumBits := checksum[0:(entrophyLen / 32)]
	entrophy = append(entrophy, checksumBits...)

	wordLen := len(entrophy) / 11
	mnemonic := []string{}
	for i := 0; i < wordLen; i++ {

		indexBytes := entrophy[(i * 11):(i + 1*11)]
		index := 0
		word := EnglishWordList[index]
	}
}

// generates a random entropy of entrophyLen bytes
func generateEntrophy(lenInBytes int) []byte {

	buffer := make([]byte, 256)

	// everyone says it's not safe
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < lenInBytes; i++ {
		bufferPtr := buffer[i*8:]
		randomNumber := rand.Uint64()
		binary.BigEndian.PutUint64(bufferPtr, randomNumber)
	}

	return buffer
}

// Calculates SHA256 of data
func calculateSha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}
