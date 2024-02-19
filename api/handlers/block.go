package handlers

import (
	"blockChain/comman"
	"blockChain/models"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) InitBlock(ctx *fiber.Ctx) error {
	err := h.BlockService.InitBlock()
	if err != nil {
		log.Println("the error: ", err)
		return err
	}
	return nil
}

func (h *Handler) AddBlock(c *fiber.Ctx) error {
	log.Println("Inside the Handler function")
	var requestBody models.AddBlockRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	log.Println("the requestBody: ", &requestBody)
	if requestBody.Data == "" {
		log.Println("Error in checking the lenght of request")
		return c.Status(fiber.ErrBadRequest.Code).JSON(comman.Failed(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("data field can't be empty").Error(),
		}))
	}
	resp, err := h.BlockService.AddBlock(&requestBody)
	if err != nil {
		log.Println("error from the service file :" + err.Error())
		return c.Status(fiber.ErrBadRequest.Code).JSON(comman.Failed(&fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: errors.New("error while executing the service").Error() + err.Error(),
		}))
	}
	return c.JSON(comman.Success(resp))
}

func (h *Handler) FindBlock(c *fiber.Ctx) error {
	var requestBody *models.FindBlockRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.Name == "" {
		log.Println("Error in checking the lenght of request")
		return c.Status(fiber.ErrBadRequest.Code).JSON(comman.Failed(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("name field can't be empty").Error(),
		}))
	}
	resp, err := h.BlockService.FindBlock(requestBody)
	if err != nil {
		log.Println("error from the service file :" + err.Error())
		return c.Status(fiber.ErrBadRequest.Code).JSON(comman.Failed(&fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: errors.New("error while executing the service").Error() + err.Error(),
		}))
	}
	return c.JSON(comman.Success(resp))
}
