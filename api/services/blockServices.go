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
	PrintBlockChain() (*models.PrintBlockChainResponseBody, error)
	DeserilizeData(*models.DeserilizeDataRequest) *models.DeserilizeDataResponse
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
	resp, err := bs.Block.FindBy(&models.DbBlock{
		Data: []byte(hashedName),
	})
	if err != nil {
		if err.Error() == "request not found" {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("error while fetching the data: " + err.Error())
	}
	return &models.FindBlockResponse{
		Response:    true,
		BlockResult: *resp,
	}, nil
}

func (bs *BlockService) PrintBlockChain() (*models.PrintBlockChainResponseBody, error) {
	err := bs.SetupDalInstance()
	if err != nil {
		return nil, err
	}
	resp, err := bs.Block.FindAll()
	if err != nil {
		return nil, errors.New("error while finding all the details of blockchain" + err.Error())
	}
	log.Println("The response = ", resp)
	var value []*models.BlockChain
	for _, items := range resp {
		blocks := &models.BlockChain{
			PrevHash: items.PrevHash,
			Data:     string(items.Data),
			Hash:     items.Hash,
		}
		value = append(value, blocks)
	}
	return &models.PrintBlockChainResponseBody{
		BlockChain: value,
	}, nil
}

func (bs *BlockService) DeserilizeData(requestBody *models.DeserilizeDataRequest) *models.DeserilizeDataResponse {
	deserlizedString := block.DeserilizeData(requestBody.Data)
	return &models.DeserilizeDataResponse{
		Data: deserlizedString,
	}
}
