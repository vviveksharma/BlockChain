package dal

import (
	"blockChain/db"
	"blockChain/models"
	"log"
)

func NewDalRequest() (Block, error) {
	return &BlockImp{}, nil
}

type Block interface {
	Create(*models.DbBlock) error
	FindAll() ([]*models.DbBlock, error)
	PreviousHash() (string, error)
}

type BlockImp struct{}

func (dal *BlockImp) Create(value *models.DbBlock) error {
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	state := transaction.Create(&value)
	if state.Error != nil {
		return state.Error
	}
	transaction.Commit()
	return nil
}

func (dal *BlockImp) FindAll() ([]*models.DbBlock, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var response []*models.DbBlock
	defer transaction.Rollback()
	result := transaction.Find(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (dal *BlockImp) PreviousHash() (string, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return "", err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return "", transaction.Error
	}
	defer transaction.Rollback()
	var response string
	if err := transaction.Model(&models.DbBlock{}).Order("id desc").Limit(1).Pluck("hash", &response).Error; err != nil {
		log.Fatal("the error in the calling the model", err)
	}
	return response, nil
}
