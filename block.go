package main

import(
        "crypto/sha256"
        "fmt"
        "time"
)

type Block struct {
        // Index           uint64
        TimeStamp       int64
        Data            []uint8
        Hash            []uint8
        PreviousHash    []uint8
}

func NewBlock(data string, previousHash []uint8) *Block {
        b := &Block{TimeStamp: time.Now().Unix(), Data: []uint8(data), PreviousHash: []uint8(previousHash)}
        // b.HashBlock()
        return b
}

func (block *Block) HashBlock() {
        hash := sha256.Sum256(block.Data)
        block.Hash = []uint8(hash[:])
}

func (block *Block) PrintBlock() {
        fmt.Printf("%s\n", block.Data)
        fmt.Printf("%x\n", block.Hash)
        fmt.Printf("%x\n\n", block.PreviousHash)
}
