package models

type AddBlockRequest struct {
	Data string `json:"data"`
}

type FindBlockRequest struct {
	Name string `json:"name"`
}
