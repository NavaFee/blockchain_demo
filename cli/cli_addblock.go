package cli

import (
	"blockchain/pbcc"
	"fmt"
	"os"
)

func (cli *CLI) addBlock(data string) {
	bc := pbcc.GetBlockchainObject()
	if bc == nil {
		fmt.Println("没有创世区块，无法添加")
		os.Exit(1)
	}
	defer bc.DB.Close()
	bc.AddBlockToBlockChain(data)

}
