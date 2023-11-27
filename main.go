package main

import (
	"fmt"
	"github.com/Xwj-art/blc/BLC"
)

func main() {
	//block := BLC.NewBlock(1, nil, []byte("the first testing"))
	//fmt.Printf("The first testing: %v\n", block)
	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	blockChain.AddBlock(
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash,
		[]byte("add Block in BlockChain"),
	)
	blockChain.AddBlock(
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash,
		[]byte("add Block in BlockChain"),
	)
	for _, block := range blockChain.Blocks {
		fmt.Printf("currentBlockHash : %x\t", block.Hash)
		fmt.Printf("prevBlockHash : %x\n", block.PrevBlockHash)
	}
	fmt.Printf("init BlockChain, the second testing: %v\n", blockChain)
}
