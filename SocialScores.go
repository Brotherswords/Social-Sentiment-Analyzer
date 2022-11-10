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
			fmt.Println(scanner.Text(), ":", val,",", totalScore)

		}
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
	fmt.Println(fileName, "sentiment score :", totalScore)
	fmt.Print(fileName, " star count: ")
	if totalScore < -5{
		fmt.Printf("1 Star")
	}else if totalScore < -1{
		fmt.Printf("2 Stars")
	}else if totalScore < 1{
		fmt.Printf("3 Stars")
	}else if totalScore < 5{
		fmt.Printf("4 Stars")
	}else if totalScore > 5{
		fmt.Printf("5 Stars")
	}
}



