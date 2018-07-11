package main

import(
        "fmt"
        "math"
)

const validateBits = 32

func main() {
        fmt.Printf("Starting go blockchain\n")

        blockchain := NewBlockchain()

        if blockchain != nil {
                fmt.Printf("Created a blockchain\n")
        }

        blockchain.AddBlock("TestBlock")

        blockchain.PrintChain()
}

func Work(block *Block) {
        // TODO
        fmt.Printf("Working for data: %s\n", block.Data)
        counter := 0

        for !Validate(counter, block.PreviousHash) {
                counter++
        }
}

func Validate(hash []uint8) {
        // TODO

}
