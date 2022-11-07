package main

import (
	//"bufio"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	m := map[string]float64{}
	totalScore := 0.0

	fmt.Println("input text:")
    scannerLocal := bufio.NewScanner(os.Stdin)
    scannerLocal.Scan()
    err := scannerLocal.Err()
    if err != nil {
        log.Fatal(err)
    }
    fileName := scannerLocal.Text()

	fmt.Println(":scream: OMG the SCORE for this MESSAGE is: ");
	csvfile, err := os.Open("socialsent.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	// There should be only one.

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
	f, err := os.Open(fileName)
 
    if err != nil {
		fmt.Println(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {

        //fmt.Println(scanner.Text())
		if val, ok := m[scanner.Text()]; ok {
			totalScore += val	
		}
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
	fmt.Println(totalScore)
}



