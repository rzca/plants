package zipcodes

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func LoadZipcodes(zipcodeDataPath string) (*ZipcodeInfoDataset, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("working directory: %s", wd)

	dat, err := os.ReadFile(zipcodeDataPath)
	log.Printf("path of data file: %s", zipcodeDataPath)
	if err != nil {
		log.Fatal(err)
	}

	return loadZipcodesInner(dat)
}

func loadZipcodesInner(bytes []byte) (*ZipcodeInfoDataset, error) {
	zipcodeInfoRows := []*ZipcodeInfoRow{}
	zipcodeMap := make(map[string]*ZipcodeInfoRow)

	r := csv.NewReader(strings.NewReader(string(bytes)))

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		log.Println(row[0], row[1], row[2], row[3])

		newZipocodeInfoRow := &ZipcodeInfoRow{
			Zipcode:   row[0],
			Zone:      row[1],
			Trange:    row[2],
			ZoneTitle: row[3],
		}
		zipcodeInfoRows = append(zipcodeInfoRows, newZipocodeInfoRow)
		zipcodeMap[newZipocodeInfoRow.Zipcode] = newZipocodeInfoRow
		if row == nil {
			log.Fatalln("exiting")
		}
	}

	zipcodeInfoDataset := &ZipcodeInfoDataset{
		InfoByZipcode: zipcodeMap,
		Count:         len(zipcodeInfoRows),
	}

	return zipcodeInfoDataset, nil
}

func (zipcodeInfoDataset *ZipcodeInfoDataset) GetZoneForZip(zipcode string) *ZipcodeInfoRow {
	val := zipcodeInfoDataset.InfoByZipcode[zipcode]

	return val

}

type ZipcodeInfoDataset struct {
	InfoByZipcode map[string]*ZipcodeInfoRow
	Count         int
}

type ZipcodeInfoRow struct {
	Zipcode   string
	Zone      string
	Trange    string
	ZoneTitle string
}
