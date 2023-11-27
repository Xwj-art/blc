package BLC

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(height int64, prevBlockHash, data []byte) {
	newBlock := NewBlock(height, prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 创建创世区块

func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)
}

// 创建区块链

func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock([]byte("init blockChain"))
	return &BlockChain{
		[]*Block{genesisBlock},
	}
}
