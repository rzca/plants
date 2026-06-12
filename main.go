package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("hello world")

	pathToZipcodeData := "data/phzm_us_zipcode_2023.csv"
	// pathToZipcodeData := "testdata/zipcodes.csv"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("working directory: %s", wd)

	path := filepath.Join(wd, pathToZipcodeData)
	dat, err := os.ReadFile(path)
	log.Printf("path of data file: %s", path)
	if err != nil {
		log.Fatal(err)
	}

	zipcodeInfoRows := []*ZipcodeInfoRow{}
	zipcodeMap := make(map[string]ZipcodeInfoRow)

	r := csv.NewReader(strings.NewReader(string(dat)))

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(row[0], row[1], row[2], row[3])

		newZipocodeInfoRow := &ZipcodeInfoRow{
			Zipcode:   row[0],
			Zone:      row[1],
			Trange:    row[2],
			ZoneTitle: row[3],
		}
		zipcodeInfoRows = append(zipcodeInfoRows, newZipocodeInfoRow)
		zipcodeMap[newZipocodeInfoRow.Zipcode] = *newZipocodeInfoRow
		if row == nil {
			log.Fatalln("exiting")
		}
	}

	zipcodeInfoDataset := &ZipcodeInfoDataset{
		InfoByZipcode: zipcodeMap,
		Count:         len(zipcodeInfoRows),
	}

	log.Printf("done, with %s rows", fmt.Sprint(len(zipcodeInfoRows)))

	log.Printf("sample for 22046: %s", zipcodeInfoDataset.GetZoneForZip("22101").Zone)

	log.Print("exiting")
}

func (zipcodeInfoDataset *ZipcodeInfoDataset) GetZoneForZip(zipcode string) ZipcodeInfoRow {
	val := zipcodeInfoDataset.InfoByZipcode[zipcode]

	return val

}

type ZipcodeInfoDataset struct {
	InfoByZipcode map[string]ZipcodeInfoRow
	Count         int
}

type ZipcodeInfoRow struct {
	Zipcode   string
	Zone      string
	Trange    string
	ZoneTitle string
}

// api endpoints should look like this
// /{dataset-id}/zipcode/{zipcode} => {}
