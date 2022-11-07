package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	m := map[string]float64{}

	m["Among Us"] = 1
	
	for key, value := range m {
		fmt.Println(key, value)
	}
	csvfile, err := os.Open("socialsent.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		wordVector, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		floatVal, err := strconv.ParseFloat(wordVector[1],32)
		m[wordVector[0]] = floatVal
	}

	for {    
		content, err := ioutil.ReadFile("Nixon.txt")

		if err != nil {
			log.Fatal(err)
		}
	}
}
	/*
	for key, value := range m {
		fmt.Println(key, value)
	}
	*/



