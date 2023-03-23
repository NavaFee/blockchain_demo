package pbcc

// 创建区块链
type BlockChain struct {
	Blocks []*Block
}

// 创建区块链，带有创世区块
func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
	genesisBlock := CreateGenesisBlock(data)
	//返回区块链对象
	return &BlockChain{[]*Block{genesisBlock}}
}
func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
	//创建新区块
	newBlock := NewBlock(data, prevHash, height)
	//将区块加入到切片中
	bc.Blocks = append(bc.Blocks, newBlock)
}
