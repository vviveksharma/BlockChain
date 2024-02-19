package block

import (
	"blockChain/dal"
	"blockChain/models"
	"errors"
	"log"
)

func (bs *Blocks) SetupDalInstance() error {
	var err error
	bs.Block, err = dal.NewDalRequest()
	if err != nil {
		return errors.New("error in the repo intilization" + err.Error())
	}
	return nil
}

func (bl *Blocks) CreateBlock(data string, prevHash string) (*models.DbBlock, error) {
	err := bl.SetupDalInstance()
	if err != nil {
		log.Println("error in setting up repo layer: ", err)
		return nil, err
	}
	block := &models.DbBlock{
		PrevHash: prevHash,
		Data:     []byte(data),
		Hash:     prevHash,
		Nonce:    0,
	}
	minedBlock, err := MineBlock(block, 4)
	if err != nil {
		log.Println("Error while mining the block:", err)
		return nil, err
	}
	hashedData := SerilizeData(minedBlock.Data)
	minedBlock.Data = []byte(hashedData)
	err = bl.Block.Create(minedBlock)
	if err != nil {
		log.Println("Error in creating the Block in the DB:", err)
		return nil, err
	}
	return minedBlock, nil
}

func (bl *Blocks) AddBlock(data string) (*models.DbBlock, error) {
	err := bl.SetupDalInstance()
	if err != nil {
		return nil, errors.New("error while setting up the DAL Layer")
	}
	response, err := bl.Block.PreviousHash()
	if err != nil {
		return nil, errors.New("error from the  DAL Layer" + err.Error())
	}
	new, err := bl.CreateBlock(data, response)
	if err != nil {
		return nil, errors.New("error in creating the new block via create block: " + err.Error())
	}
	return new, nil
}

func (bl *Blocks) Genesis() *models.DbBlock {
	resp, err := bl.CreateBlock("Genesis", "")
	if err != nil {
		log.Println("Error while calling the genesis block")
		return nil
	}
	return resp
}

func (bl *Blocks) InitBlockChain() *BlockChain {
	block := bl.Genesis()
	bl.Chain.Blocks = append(bl.Chain.Blocks, block)
	return &BlockChain{Blocks: bl.Chain.Blocks}
}
