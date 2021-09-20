package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"massivleads/interfaces"
	"massivleads/interfaces/middlewares"

	"github.com/joho/godotenv"
	"massivleads/exceptions"
	"massivleads/logger"
)

var _ = godotenv.Load()

func main() {
	logger.Init()
	procSize, err := strconv.ParseInt(os.Getenv("PROCS"), 10, 0)
	if err != nil {
		procSize = int64(runtime.GOMAXPROCS(-1))
	}
	runtime.GOMAXPROCS(int(procSize))

	app := fiber.New(
		fiber.Config{
			CaseSensitive: true,
			ServerHeader:  os.Getenv("SERVER_HEADER"),
			Prefork:       os.Getenv("PREFORK") == "yes",
			ErrorHandler:  exceptions.GlobalExceptionHandler,
		},
	)

	// Initialize middlewares
	middlewares.SetMiddlewares(app)

	// Initialize routes
	interfaces.CreateRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
