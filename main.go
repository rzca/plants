package main

import (
	"fmt"
	"log"
	"plants/zipcodes"
)

func main() {
	fmt.Println("hello world")

	pathToZipcodeData := "data/phzm_us_zipcode_2023.csv"
	// pathToZipcodeData := "testdata/zipcodes.csv"

	zipcodeInfoDataset, err := zipcodes.LoadZipcodes(pathToZipcodeData)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("done, with %s rows", fmt.Sprint(len(zipcodeInfoDataset.InfoByZipcode)))

	log.Printf("sample for 22046: %s", zipcodeInfoDataset.GetZoneForZip("22101").Zone)

	log.Print("exiting")
}

// api endpoints should look like this
// /{dataset-id}/zipcode/{zipcode} => {}
