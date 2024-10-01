package models

import "github.com/google/uuid"

type AddBlockRespose struct {
	Id uuid.UUID `json:"id"`
}

type FindBlockResponse struct {
	Response    bool    `json:"response"`
	BlockResult DbBlock `json:"blockResult"`
}

type BlockChain struct {
	PrevHash string `json:"prev_hash"`
	Data     string `json:"data"`
	Hash     string `json:"hash"`
}

type PrintBlockChainResponseBody struct {
	BlockChain []*BlockChain `json:"blockChain"`
}

type DeserilizeDataResponse struct {
	Data string `json:"data"`
}
