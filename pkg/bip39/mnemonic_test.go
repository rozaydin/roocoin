package bip39

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testVectors struct {
	English []testUnit `json:"english"`
}

type testUnit struct {
	Entropy    string `json:"entropy"`
	Mnemonic   string `json:"mnemonic"`
	Seed       string `json:"seed"`
	PrivateKey string `json:"privatekey"`
}

var (
	vectors testVectors
)

func init() {
	err := loadTestVector("./vectors.json")
	fmt.Printf("test vector loaded!")
	if err != nil {
		panic(fmt.Errorf("failed to load vector data for testing"))
	}
}

// utils

func loadTestVector(relativePath string) error {

	data, err := ioutil.ReadFile(relativePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &vectors)
	if err != nil {
		return err
	}

	fmt.Printf("vectors: %v+\n", vectors)

	return nil
}

func TestVector(t *testing.T) {
	for _, unit := range vectors.English {		
		entropy, _ := hex.DecodeString(unit.Entropy)
		mnemonic := generateMnemonic(entropy)
		assert.Equal(t, strings.Split(unit.Mnemonic, " "), mnemonic)
	}
}

func TestReadBit(t *testing.T) {

	// 32 bits 1111 1111 0000 0000 1111 1111 0000 0000
	testData := []byte{0xFF, 0x00, 0xFF, 0x00}
	for i := 0; i < 8; i++ {
		readBit := readBit(i, testData)
		assert.Equal(t, true, readBit)
	}

	for i := 8; i < 16; i++ {
		readBit := readBit(i, testData)
		assert.Equal(t, false, readBit)
	}

	for i := 16; i < 24; i++ {
		readBit := readBit(i, testData)
		assert.Equal(t, true, readBit)
	}

	for i := 24; i < 32; i++ {
		readBit := readBit(i, testData)
		assert.Equal(t, false, readBit)
	}

}

func TestReadIndex(t *testing.T) {
	// 32 bits 1111 1111 0000 0000 1111 1111 0000 0000
	testData := []byte{0xFF, 0x00, 0xFF, 0x00}
	byte0 := readIndex(0, 8, testData)
	assert.Equal(t, uint16(0xFF), byte0)

	byte1 := readIndex(8, 8, testData)
	assert.Equal(t, uint16(0x00), byte1)

	byte2 := readIndex(16, 8, testData)
	assert.Equal(t, uint16(0xFF), byte2)

	byte3 := readIndex(24, 8, testData)
	assert.Equal(t, uint16(0x00), byte3)

	testData = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	indexCount := (len(testData) * 8) / 11

	for i := 0; i < indexCount; i++ {
		readIndex := readIndex(i*11, 11, testData)
		assert.Equal(t, uint16(2047), readIndex)
	}
}
