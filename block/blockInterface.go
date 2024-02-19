package block

import (
	"blockChain/dal"
	"blockChain/models"
)

type BlockChain struct {
	Blocks []*models.DbBlock
}

type Blocks struct {
	Block dal.Block
	Chain BlockChain
}

type IBlocks interface {
	InitBlockChain() *BlockChain
	CreateBlock(data string, prevHash string) (*models.DbBlock, error)
	AddBlock(data string) (*models.DbBlock, error)
	Genesis() *models.DbBlock
}

func NewBlocksService() (IBlocks, error) {
	return &Blocks{}, nil
}
