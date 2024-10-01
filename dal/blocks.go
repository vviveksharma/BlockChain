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
	FindBy(conditions *models.DbBlock) (*models.DbBlock, error)
}

type BlockImp struct{}

func (dal *BlockImp) Create(value *models.DbBlock) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return nil
	}
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
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return nil, err
	}
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
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return "", err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return "", err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return "", transaction.Error
	}
	defer transaction.Rollback()
	var resp models.DbBlock
	err = transaction.Model(&models.DbBlock{}).Order("created_at DESC").Limit(1).Find(&resp).Error
	if err != nil {
		return "", err
	}
	transaction.Commit()
	return resp.Hash, nil
}

func (dal *BlockImp) FindBy(conditions *models.DbBlock) (*models.DbBlock, error) {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return nil, err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var response *models.DbBlock
	resp := transaction.Last(&response, conditions)
	if resp.Error != nil {
		return nil, resp.Error
	}
	transaction.Commit()
	return response, nil
}