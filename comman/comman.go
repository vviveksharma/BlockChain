package comman

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
)

func Getenv() error {
	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file" + err.Error())
		return err
	}
	return nil
}

type ResponseBody struct {
	Status string       `json:"status"`
	Result interface{}  `json:"result"`
	Error  *fiber.Error `json:"error,omitempty"`
}

func Success(response interface{}) *ResponseBody {
	return &ResponseBody{SUCCESS, response, nil}
}

func Failed(err *fiber.Error) *ResponseBody {
	return &ResponseBody{FAILED, nil, err}
}
