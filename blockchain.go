package main

type Blockchain struct {
        chain []*Block
}

func NewBlockchain() *Blockchain {
        bc := new(Blockchain)

        // set genesis block
        genBlock := NewBlock("Genesis Block", []uint8{})
        bc.chain = append(bc.chain, genBlock)

        return bc
}

func (bc *Blockchain) AddBlock(data string) {
        previous := bc.chain[len(bc.chain) - 1]
        newBlock := NewBlock(data, previous.Hash)

        // work for the block
        hash = Work(newBlock)

        bc.chain = append(bc.chain, newBlock)
}

func (bc *Blockchain) PrintChain() {
        for _, block := range bc.chain {
                block.PrintBlock()
        }
}

func Work(block *Block) []uint8 {

}
