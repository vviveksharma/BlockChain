package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbBlock struct {
	Id         uuid.UUID `gorm:"primaryKey,column:id"`
	Created_At time.Time `gorm:"column:created_at;not null"`
	PrevHash   string    `gorm:"column:prev_hash;type:varchar(500)"`
	Data       []byte    `gorm:"column:data;type:varchar(500);not null"`
	Hash       string    `gorm:"column:hash;type:varchar(500);not null"`
	Nonce      int       `gorm:"column:nonce;not null"`
}

func (DbBlock) TableName() string {
	return "block_tbl"
}

func (*DbBlock) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}
