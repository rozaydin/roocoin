package bip39

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

func generateMnemonic(entropy []byte) []string {

	entropyBitLen := len(entropy) * 8
	checksumBitLen := entropyBitLen / 32
	checksum := calculateSha256(entropy)
	checksumBits := checksum[0]
	entropy = append(entropy, checksumBits)

	wordLen := (entropyBitLen + checksumBitLen) / 11
	mnemonic := []string{}
	for i := 0; i < wordLen; i++ {
		mnemonic = append(mnemonic, calculateWord(i, entropy))
	}

	return mnemonic
}

func GeneratePrivateKey(keyLength int) (string, []byte) {
	entropy := generateEntrophy(keyLength / 8)
	mnemonics := generateMnemonic(entropy)
	return strings.Join(mnemonics, " "), StretchKey(mnemonics)
}

func StretchKey(mnemonics []string) []byte {
	salt := "mnemonic"
	password := strings.Join(mnemonics, "")
	return pbkdf2.Key([]byte(password), []byte(salt), 2048, 512, sha512.New512_256)
}

// generates a random entropy of entrophyLen bytes
func generateEntrophy(lenInBytes int) []byte {

	buffer := make([]byte, lenInBytes)

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

func calculateWord(wordIndex int, entropy []byte) string {
	word := EnglishWordList[readIndex(wordIndex*11, 11, entropy)]
	return word
}

func readIndex(bitIndex, bitCount int, data []byte) uint16 {
	var index uint16 = 0
	for i := 0; i < bitCount; i++ {
		if readBit(bitIndex+i, data) {
			index = index | 0x0001
		}

		if i != (bitCount - 1) {
			// shift 1 bit
			index = index << 1
		}
	}

	return index & 0x07FF
}

func readBit(bitIndex int, data []byte) bool {
	byteIndex := bitIndex / 8
	// flip offset and only take 4 LSB
	// makes 6 -> 1, 7 -> 0
	offset := (^(bitIndex % 8)) & 0x07

	_byte := data[byteIndex]
	return _byte>>offset&0x01 == 0x01
}
