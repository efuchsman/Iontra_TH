package main

import (
	"sync"

	"github.com/efuchsman/Iontra_TH/internal/citiesapi"
	log "github.com/sirupsen/logrus"
)

func main() {
	citiesapi := citiesapi.NewClient()
	validElements, err := citiesapi.ReturnValidTmpElements()
	if err != nil {
		log.Fatalf("Error returing valid tmp elements: %v", err)
	}
	invalidElements, err := citiesapi.ReturnInvalidTmpElements()
	if err != nil {
		log.Fatalf("Error returning invalid tmp elements: %v", err)
	}
	badFiles := citiesapi.ReturnUnprocessableFiles()

	// Concurrent processing through the use of goroutines to executes the build of the
	// two JSON files and the CSV files and waits to check for errors until all concurrent goroutines are finished
	var wg sync.WaitGroup

	// Channel for catching errors within the goroutines
	errChan := make(chan error, 3)
	wg.Add(3)
	go func() {
		defer wg.Done()
		log.Infof("Creating valid elements JSON")
		if err := citiesapi.CreateValidElementsJSON(validElements, "valid_elements.json"); err != nil {
			errChan <- err
		}
		log.Infof("Finished creating valid elements JSON")
	}()
	go func() {
		defer wg.Done()
		log.Infof("Creating invalid elements JSON")
		if err := citiesapi.CreateInvalidElementsJSON(invalidElements, "invalid_elements.json"); err != nil {
			errChan <- err
		}
		log.Infof("Finished creating invalid elements JSON")
	}()
	go func() {
		defer wg.Done()
		log.Infof("Creating unprocessable files CSV")
		if err := citiesapi.CreateUnprocessableFilesCSV(badFiles.Files, "unprocessable_files.csv"); err != nil {
			errChan <- err
		}
		log.Infof("Finished creating unprocessable files CSV")
	}()

	// Waits for goroutines to execute and then closes the error channel
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Check for errors caught within the goroutines
	for err := range errChan {
		log.Fatalf("Received error from goroutine: %v", err)
	}
}
