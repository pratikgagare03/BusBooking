package main

import (
	"busbooking/availability"
	"busbooking/book"
	"busbooking/logger"
	"busbooking/types"
	"busbooking/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func setupRoutes(buses *[]types.Bus) {
	http.HandleFunc("/check_availability", availability.Check_availability(buses))
	http.HandleFunc("/book", book.BookingHandler(buses))
}

func startServer() {
	fmt.Printf("localhost started at port%v", os.Getenv("APP_PORT"))
	err := http.ListenAndServe(os.Getenv("APP_PORT"), nil)
	if err != nil {
		logger.Logs.Error().Err(err)
		return
	}
}

func main() {
	logger.Logs.Info().Msg("Started Main")
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logs.Error().Err(err)
	}
	defer logger.File.Close()
	buses := utils.CreateBuses()
	setupRoutes(buses)
	startServer()

	logger.Logs.Info().Msg("Main Function over")
}
