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
	Blocks    []*Block
	Consensus *Consensus
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

	if bc.Consensus.ProposeBlock(newBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		fmt.Println("Block rejected by consensus")
	}
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
	// Create the blockchain with consensus
	blockchain := Blockchain{
		Blocks:    []*Block{CreateGenesisBlock()},
		Consensus: NewConsensus(6), // Example with 6 nodes
	}

	// Simulate multiple nodes adding blocks
	nodeData := []struct {
		NodeID     int
		BlockData  string
		Difficulty int
	}{
		{NodeID: 1, BlockData: "I just woke up", Difficulty: 4},
		{NodeID: 2, BlockData: "Kevin proposed a strong idea", Difficulty: 6},
		{NodeID: 6, BlockData: "Third Block", Difficulty: 4},
		{NodeID: 4, BlockData: "Fourth Block", Difficulty: 2},
		{NodeID: 5, BlockData: "I am alarmed by the spending", Difficulty: 4},
		{NodeID: 2, BlockData: "Sixth Block", Difficulty: 1},
		{NodeID: 3, BlockData: "Seventh Block", Difficulty: 6},
		{NodeID: 4, BlockData: "Eighth Block", Difficulty: 1},
		{NodeID: 1, BlockData: "Ninth Block", Difficulty: 3},
		{NodeID: 2, BlockData: "Tenth Block", Difficulty: 5},
		{NodeID: 5, BlockData: "Eleventh Block", Difficulty: 2},
		{NodeID: 3, BlockData: "Twelfth Block", Difficulty: 3},
		{NodeID: 1, BlockData: "Thirteenth Block", Difficulty: 4},
		{NodeID: 6, BlockData: "Fourteenth Block", Difficulty: 1},
		{NodeID: 4, BlockData: "Fifteenth Block", Difficulty: 6},
		{NodeID: 1, BlockData: "I just woke up", Difficulty: 4},
		{NodeID: 2, BlockData: "Kevin proposed a strong idea", Difficulty: 6},
		{NodeID: 6, BlockData: "Third Block", Difficulty: 4},
		{NodeID: 4, BlockData: "Fourth Block", Difficulty: 2},
		{NodeID: 5, BlockData: "I am alarmed by the spending", Difficulty: 4},
		{NodeID: 2, BlockData: "Sixth Block", Difficulty: 1},
		{NodeID: 3, BlockData: "Seventh Block", Difficulty: 6},
		{NodeID: 4, BlockData: "Eighth Block", Difficulty: 1},
		{NodeID: 1, BlockData: "Ninth Block", Difficulty: 3},
		{NodeID: 2, BlockData: "Tenth Block", Difficulty: 5},
		{NodeID: 5, BlockData: "Eleventh Block", Difficulty: 2},
		{NodeID: 3, BlockData: "Twelfth Block", Difficulty: 3},
		{NodeID: 1, BlockData: "Thirteenth Block", Difficulty: 4},
		{NodeID: 6, BlockData: "Fourteenth Block", Difficulty: 1},
		{NodeID: 4, BlockData: "Fifteenth Block", Difficulty: 6},
		{NodeID: 2, BlockData: "Sixteenth Block", Difficulty: 2},
		{NodeID: 4, BlockData: "Seventeenth Block", Difficulty: 3},
		{NodeID: 5, BlockData: "Eighteenth Block", Difficulty: 5},
		{NodeID: 6, BlockData: "Nineteenth Block", Difficulty: 4},
		{NodeID: 5, BlockData: "Twentieth Block", Difficulty: 1},
		{NodeID: 2, BlockData: "Sixteenth Block", Difficulty: 2},
		{NodeID: 4, BlockData: "Seventeenth Block", Difficulty: 3},
		{NodeID: 5, BlockData: "Eighteenth Block", Difficulty: 5},
		{NodeID: 6, BlockData: "Nineteenth Block", Difficulty: 4},
		{NodeID: 5, BlockData: "Twentieth Block", Difficulty: 1},
	}

	for _, data := range nodeData {
		fmt.Printf("Node %d proposing block...\n", data.NodeID)
		blockchain.AddBlock(data.BlockData, data.Difficulty)
	}

	blockchain.PrintBlockchain()
}
