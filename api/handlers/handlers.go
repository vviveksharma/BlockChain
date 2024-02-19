package handlers

import (
	"blockChain/api/services"
	"log"
)

type Handler struct {
	Logger       *log.Logger
	BlockService services.IBlockService
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{Logger: logger}
}

func (h *Handler) BlockServiceInstance(bs services.IBlockService) *Handler {
	h.BlockService = bs
	return h
}
