package main

type Blockchain struct {
        chain []Block
}

func NewBlockchain() *Blockchain {
        return &Blockchain{}
}
