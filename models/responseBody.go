package models

import "github.com/google/uuid"

type AddBlockRespose struct {
	Id uuid.UUID `json:"id"`
}

type FindBlockResponse struct {
	Response bool `json:"response"`
}
