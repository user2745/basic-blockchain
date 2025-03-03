package main

import (
	"fmt"
)

type Node struct {
	ID    int
	Block *Block
}

type Consensus struct {
	Nodes []*Node
}

func NewConsensus(nodeCount int) *Consensus {
	nodes := make([]*Node, nodeCount)
	for i := 0; i < nodeCount; i++ {
		nodes[i] = &Node{ID: i}
	}
	return &Consensus{Nodes: nodes}
}

func (c *Consensus) ProposeBlock(block *Block) bool {
	// Simulate the pBFT consensus process
	agreeCount := 0
	for _, node := range c.Nodes {
		if node.ValidateBlock(block) {
			agreeCount++
		}
	}

	// Assuming more than 2/3 agreement is required
	if agreeCount > (2 * len(c.Nodes) / 3) {
		fmt.Println("Block accepted by consensus")
		return true
	}

	fmt.Println("Block rejected by consensus")
	return false
}

func (n *Node) ValidateBlock(block *Block) bool {
	// Simulate block validation
	return block.CalculateHash() == block.Hash
}

// pBFT consensus mechanism
// Reference: https://www.geeksforgeeks.org/practical-byzantine-fault-tolerancepbft/
