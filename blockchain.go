package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func (b *Block) CalculateHash() string {
	record := fmt.Sprintf("%d%s%s%s%d", b.Index, b.Timestamp, b.Data, b.PrevHash, b.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (b *Block) MineBlock(difficulty int) {

	// Validate difficulty
	if difficulty > 64 {
		panic("Difficulty too high!  Maximum allowed is 64")
	}

	target := ""
	for i := 0; i < difficulty; i++ {
		target += "0"
	}

	// Initialize the hash
	b.Hash = b.CalculateHash()

	// Mine the block
	for b.Hash[:difficulty] != target {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}

	fmt.Println("Block mined: ", b.Hash)
}

type Blockchain struct {
	Blocks []*Block
}

func CreateGenesisBlock() *Block {

	block := &Block{

		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}

	block.Hash = block.CalculateHash()

	return block
}

func (bc *Blockchain) AddBlock(data string, difficulty int) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:     len(bc.Blocks),
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}

	newBlock.MineBlock(difficulty)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Println("--------------------------------")
	}
}

func main() {
	// Create the blockchain
	blockchain := Blockchain{Blocks: []*Block{CreateGenesisBlock()}}

	// Add Blocks
	blockchain.AddBlock("First Block", 4)
	blockchain.AddBlock("Second Block", 4)
	blockchain.AddBlock("Third Block", 4)

	blockchain.PrintBlockchain()

}
