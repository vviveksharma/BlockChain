package services

import (
	"blockChain/block"
	"blockChain/dal"
	"blockChain/models"
	"errors"
	"log"
)

type BlockService struct {
	Block  dal.Block
	Blocks block.IBlocks
}

type IBlockService interface {
	InitBlock() error
	AddBlock(*models.AddBlockRequest) (*models.AddBlockRespose, error)
	FindBlock(requestBody *models.FindBlockRequest) (*models.FindBlockResponse, error)
}

func NewBlockService() *BlockService {
	return &BlockService{}
}

func (bs *BlockService) SetupDalInstance() error {
	var err error
	bs.Block, err = dal.NewDalRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

func (bs *BlockService) SetupBlockInstance() error {
	var err error
	bs.Blocks, err = block.NewBlocksService()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

func (bs *BlockService) InitBlock() error {
	log.Println("Inside the service implementation *************")
	res := bs.Blocks.InitBlockChain()
	log.Println("The result from the initBlockChain: ", res)
	return nil
}

func (bs *BlockService) AddBlock(requestBody *models.AddBlockRequest) (*models.AddBlockRespose, error) {
	err := bs.SetupBlockInstance()
	if err != nil {
		return nil, errors.New("error in setting up the BlockService: " + err.Error())
	}
	err = bs.SetupDalInstance()
	if err != nil {
		return nil, errors.New("error in setting up the Dal Service in the AddBlock: " + err.Error())
	}
	resp, err := bs.Blocks.AddBlock(requestBody.Data)
	if err != nil {
		log.Println("The service addblock err: " + err.Error())
		return nil, errors.New("error while adding the block via AddBlock" + err.Error())
	}

	return &models.AddBlockRespose{
		Id: resp.Id,
	}, nil
}

func (bs *BlockService) FindBlock(requestBody *models.FindBlockRequest) (*models.FindBlockResponse, error) {
	err := bs.SetupDalInstance()
	if err != nil {
		return nil, errors.New("error in setting up the dal layer instance: " + err.Error())
	}
	hashedName := block.SerilizeData([]byte(requestBody.Name))
	resp, err := bs.Block.FindByName(&models.DbBlock{
		Data: []byte(hashedName),
	})
	if err != nil {
		if err.Error() == "request not found" {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("error while fetching the data: " + err.Error())
	}
	log.Println(resp.Data)
	log.Println(resp.Id)
	return &models.FindBlockResponse{
		Response: true,
	}, nil
}
