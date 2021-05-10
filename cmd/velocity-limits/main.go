package main

import (
	"log"
	"time"
	"velocity-limits/config"
	"velocity-limits/internal/service"
	"velocity-limits/internal/storage"
)

func main() {
	start := time.Now()

	config := config.LoadConfig("../../config/")
	storage := storage.NewStorage()

	transactions, err := service.GetTransactionsFromInputFile(&config)
	if err != nil {
		log.Fatal("Error - from GetTransaction function: ", err)
	}

	responses := service.LoadFunds(&config, transactions, storage)

	// write to file
	if err = service.WriteResponsesToOutputFile(&config, responses); err != nil {
		log.Fatal("Error -  from WriteResponsesToOutputFile function: ", err)
	}

	elapsed := time.Since(start)
	log.Println("Total program execution time: ", elapsed)
}