// Package main is an entrypoint for our application. It reads the file, loads transactions
// and write to output file.
package main

import (
	"log"
	"velocity-limits/config"
	"velocity-limits/internal/service"
	"velocity-limits/internal/storage"
)

func main() {
	// Define project root path and get config, storage structs.
	projectRootPath := "../../"
	config := config.LoadConfig(projectRootPath + "config/")
	storage := storage.NewStorage()

	// Try reading an input file and get transactions array.
	transactions, err := service.GetTransactionsFromInputFile(&config, projectRootPath)
	if err != nil {
		log.Fatal("Error - from GetTransaction function: ", err)
	}

	// Load funds from transactions and create responses array to store result from each transaction.
	responses := service.LoadFunds(&config, transactions, storage)

	// Write responses to the output file.
	if err = service.WriteResponsesToOutputFile(&config, responses, projectRootPath); err != nil {
		log.Fatal("Error -  from WriteResponsesToOutputFile function: ", err)
	}
}
