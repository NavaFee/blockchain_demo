package cli

import "blockchain/pbcc"

func (cli *CLI) createGenesisBlockchain(data string) {
	pbcc.CreateBlockChainWithGenesisBlock(data)
}
