package main

type Block struct {
        Index           uint64
        TimeStamp       int64
        Data            []uint8
        Hash            []uint8
        PreviousHash    []uint8
}

