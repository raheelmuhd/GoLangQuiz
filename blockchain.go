package main
//this is a change im making in this file to then merge with main branch.
import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          string
	PreviousBlock *Block
	//adding hash 
    	Hash          string
}

type Blockchain struct {
	blocks []*Block
}

func NewBlock(data string, previousBlock *Block) *Block {
	return &Block{Timestamp: time.Now().Unix(), Data: data, PreviousBlock: previousBlock}
}

func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, previousBlock)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", nil)
	return &Blockchain{blocks: []*Block{genesisBlock}}
}

func (bc *Blockchain) DisplayAllBlocks() {
	for _, block := range bc.blocks {
		fmt.Printf("Timestamp: %d, Data: %s\n", block.Timestamp, block.Data)
	}
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")
	bc.DisplayAllBlocks()
}
