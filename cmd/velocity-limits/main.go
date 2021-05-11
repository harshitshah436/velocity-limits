package main

import (
	"log"
	"velocity-limits/config"
	"velocity-limits/internal/service"
	"velocity-limits/internal/storage"
)

func main() {
	projectRootPath := "../../"
	config := config.LoadConfig(projectRootPath + "config/")
	storage := storage.NewStorage()

	transactions, err := service.GetTransactionsFromInputFile(&config, projectRootPath)
	if err != nil {
		log.Fatal("Error - from GetTransaction function: ", err)
	}

	responses := service.LoadFunds(&config, transactions, storage)

	// write to file
	if err = service.WriteResponsesToOutputFile(&config, responses, projectRootPath); err != nil {
		log.Fatal("Error -  from WriteResponsesToOutputFile function: ", err)
	}
}
