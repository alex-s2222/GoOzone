package main

import (
	// "compress/bzip2"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"net/http"


	"github.com/pkg/errors"
)


func main(){
	f, err := os.Open("geoip.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(countIsoCode(f, "RU"))
}

func mainHttp(){
	res, err := http.Get("http://192.168.5.110:8000/storage/temp/geoip.csv.bz2")
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		log.Fatal("wrong status code:", res.StatusCode)
	}
	fmt.Println(countIsoCode(res.Body, "RU"))
}

func countIsoCode(reader io.Reader, country string) (int, error){
	// bzipReader := bzip2.NewReader(reader)
	csvReader := csv.NewReader(reader)
	counter := 0
	rowId := 0
	for {
		rowId ++ 
		if rowId%100_000 == 0{
			log.Println("processed", rowId, "rows")
		}
		row, err := csvReader.Read()
		if err != nil{
			if err == io.EOF{
				break
			}
			return 0, errors.Wrap(err, "reading csv")
		}
		if len(row) < 3{
			continue
		}
		if row[4] == country{
			counter++
		}
	}
	log.Printf("total rows: %d", rowId)
	return counter, nil
}


