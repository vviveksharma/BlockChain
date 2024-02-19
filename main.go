package main

import (
	routes "blockChain/api"
	"blockChain/api/handlers"
	"blockChain/api/services"
	"blockChain/block"
	"blockChain/dal"
	"blockChain/db"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		db, err  := db.NewDbRequest()
		if err != nil {
			log.Fatal("error in creating a DB request")
		}
		resp, err := db.InitDB()
		if err != nil {
			log.Println("error in starting the DataBase: ", err)
		}
		if resp != nil {
			log.Println("THE DATABASE IS RUNNING")
			dal, err := dal.NewDalRequest()
			if err != nil {
				log.Println("error in checking the BlockChain (setting instance):", err)
			}
			resp, err := dal.FindAll()
			if err != nil {
				log.Println("error in checking the BlockChain (findALL):", err)
			}
			if len(resp) == 0 {
				block, _ := block.NewBlocksService()
				_ = block.InitBlockChain()
			} else {
				log.Println("DataBase already has the genesis Node")
			}
		}
		close(done)
	}()
	var Log *log.Logger
	handlers := handlers.NewHandler(Log).BlockServiceInstance(services.NewBlockService())
	routes.Routes(app, handlers)
	err := app.Listen(":3000")
	if err != nil {
		log.Print("error in starting the server:", err)
	}
}
