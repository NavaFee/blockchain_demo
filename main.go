package main

import (
	"blockchain/pbcc"
	"fmt"
)

func main() {
	//创建带有创世区块的区块链
	blockchain := pbcc.CreateBlockChainWithGenesisBlock("i am genesisblock")
	//添加区块
	blockchain.AddBlockToBlockChain("first block",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockChain("second block",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockChain("third block",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	//循环遍历区块信息
	for _, block := range blockchain.Blocks {
		fmt.Printf("Timestamp: %d\n", block.TimeStamp)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevBlockHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("height: %d\n", block.Height)
		fmt.Println("----------------------------------------------")
	}

}
